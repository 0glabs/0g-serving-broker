package server

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	image "github.com/0glabs/0g-serving-broker/common/docker"
	"github.com/0glabs/0g-serving-broker/common/log"
	"github.com/0glabs/0g-serving-broker/common/phala"
	"github.com/0glabs/0g-serving-broker/common/token"
	"github.com/0glabs/0g-serving-broker/fine-tuning/config"
	providercontract "github.com/0glabs/0g-serving-broker/fine-tuning/internal/contract"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/ctrl"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/db"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/handler"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/services"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/storage"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

//go:generate swag fmt
//go:generate swag init --dir ./,../../ --output ../../doc

//	@title			0G Compute Network Fine-tuning Provider API
//	@version		0.1.0
//	@description	These APIs allows providers to interact with the 0G Compute Fine Tune Service
//	@host			localhost:3080
//	@BasePath		/v1
//	@in				header

func Main() {
	cfg, logger, err := initializeBaseComponents()
	if err != nil {
		panic(err)
	}

	if err := token.CheckPythonEnv(logger); err != nil {
		panic(err)
	}

	ctx := context.Background()
	imageChan := buildImageIfNeeded(ctx, cfg, logger)

	services, err := initializeServices(ctx, cfg, logger)
	if err != nil {
		panic(err)
	}
	defer services.contract.Close()

	if err := runApplication(ctx, services, logger, imageChan); err != nil {
		panic(err)
	}
}

type ApplicationServices struct {
	db            *db.DB
	storageClient *storage.Client
	contract      *providercontract.ProviderContract
	phalaService  *phala.PhalaService
	ctrl          *ctrl.Ctrl
	executor      *services.Executor
	settlement    *services.Settlement
}

func initializeBaseComponents() (*config.Config, log.Logger, error) {
	config := config.GetConfig()
	logger, err := log.GetLogger(&config.Logger)
	return config, logger, err
}

func buildImageIfNeeded(ctx context.Context, config *config.Config, logger log.Logger) chan bool {
	imageChan := make(chan bool, 1)

	if !config.Images.BuildImage {
		imageChan <- true
		close(imageChan)
		return imageChan
	}

	go func() {
		defer close(imageChan)

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			logger.Errorf("Failed to create docker client: %v", err)
			return
		}
		defer cli.Close()

		imageName := config.Images.ExecutionImageName
		buildImage := true
		if !config.Images.OverrideImage {
			exists, err := image.ImageExists(ctx, cli, imageName)
			if err != nil {
				logger.Errorf("Failed to check image existence: %v", err)
				return
			}

			logger.Debugf("Image %s status %v.", imageName, exists)
			if exists {
				buildImage = false
			}
		}

		if buildImage {
			logger.Debugf("Build image %s", imageName)
			err := image.ImageBuild(ctx, cli, "./fine-tuning/execution/transformer", imageName)
			if err != nil {
				logger.Errorf("Failed to build image: %v", err)
				return
			}

			logger.Debugf("Docker image %s built successfully!", imageName)
		}

		imageChan <- true
	}()

	return imageChan
}

func initializeServices(ctx context.Context, cfg *config.Config, logger log.Logger) (*ApplicationServices, error) {
	db, err := db.NewDB(cfg, logger)
	if err != nil {
		return nil, err
	}
	if err := db.Migrate(); err != nil {
		return nil, err
	}

	storageClient, err := storage.New(cfg, logger)
	if err != nil {
		return nil, err
	}

	contract, err := providercontract.NewProviderContract(cfg, logger)
	if err != nil {
		return nil, err
	}

	phalaClientType := phala.TEE
	if os.Getenv("NETWORK") == "hardhat" {
		phalaClientType = phala.Mock
	}

	phalaService, err := phala.NewPhalaService(phalaClientType)
	if err != nil {
		return nil, err
	}

	ctrl := ctrl.New(db, cfg, contract, phalaService, logger)

	executor, err := services.NewExecutor(db, cfg, contract, logger, storageClient, phalaService)
	if err != nil {
		return nil, err
	}

	settlement, err := services.NewSettlement(db, contract, cfg, phalaService, logger)
	if err != nil {
		return nil, err
	}

	return &ApplicationServices{
		db:            db,
		storageClient: storageClient,
		contract:      contract,
		phalaService:  phalaService,
		ctrl:          ctrl,
		executor:      executor,
		settlement:    settlement,
	}, nil
}

func runApplication(ctx context.Context, services *ApplicationServices, logger log.Logger, imageChan <-chan bool) error {
	if err := services.phalaService.SyncQuote(ctx); err != nil {
		return err
	}

	if err := services.db.MarkInProgressTasksAsFailed(); err != nil {
		return err
	}

	if err := services.executor.Start(ctx); err != nil {
		return err
	}

	engine := gin.New()
	h := handler.New(services.ctrl, logger)
	h.Register(engine)

	if _, ok := <-imageChan; !ok {
		return errors.New("image build failed")
	}

	if err := services.ctrl.SyncServices(ctx); err != nil {
		return err
	}

	if err := services.settlement.Start(ctx); err != nil {
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Listen and Serve, config port with PORT=X
	go func() {
		if err := engine.Run(); err != nil {
			logger.Errorf("HTTP server error: %v", err)
			stop <- os.Interrupt
		}
	}()

	<-stop
	logger.Info("Shutting down server...")
	return nil
}
