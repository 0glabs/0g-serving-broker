package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/0glabs/0g-serving-broker/common/log"
	"github.com/0glabs/0g-serving-broker/fine-tuning/config"
	providercontract "github.com/0glabs/0g-serving-broker/fine-tuning/internal/contract"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/ctrl"
	database "github.com/0glabs/0g-serving-broker/fine-tuning/internal/db"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/handler"
	"github.com/gin-gonic/gin"
)

//go:generate swag fmt
//go:generate swag init --dir ./,../../ --output ../../doc

//	@title			0G Serving Provider Broker API
//	@version		0.2.0
//	@description	These APIs allows customers to interact with the 0G Compute Fine Tune Service
//	@host			localhost:3080
//	@BasePath		/v1
//	@in				header

func Main() {
	config := config.GetConfig()

	logger, err := log.GetLogger(&config.Logger)
	if err != nil {
		panic(err)
	}

	db, err := database.NewDB(config, logger)
	if err != nil {
		panic(err)
	}
	if err := db.Migrate(); err != nil {
		panic(err)
	}

	contract, err := providercontract.NewProviderContract(config, logger)
	if err != nil {
		panic(err)
	}
	defer contract.Close()

	ctrl := ctrl.New(db, contract, config.Services, logger)

	ctx := context.Background()
	err = ctrl.SyncServices(ctx)
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	h := handler.New(ctrl, logger)
	h.Register(engine)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Listen and Serve, config port with PORT=X
		if err := engine.Run(); err != nil {
			panic(err)
		}
	}()

	logger.Info("Server started")
	<-stop

	if err := ctrl.DeleteAllService(ctx); err != nil {
		logger.Errorf("Error deleting all services: %v", err)
	}
}