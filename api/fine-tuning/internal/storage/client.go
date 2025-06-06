package storage

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"

	"github.com/0glabs/0g-serving-broker/common/chain"
	"github.com/0glabs/0g-serving-broker/common/errors"
	"github.com/0glabs/0g-serving-broker/common/log"
	"github.com/0glabs/0g-serving-broker/common/util"
	"github.com/0glabs/0g-serving-broker/fine-tuning/config"
	"github.com/0glabs/0g-storage-client/common"
	"github.com/0glabs/0g-storage-client/common/blockchain"
	"github.com/0glabs/0g-storage-client/core"
	"github.com/0glabs/0g-storage-client/indexer"
	"github.com/0glabs/0g-storage-client/transfer"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/sirupsen/logrus"
)

var nRetriesToUpload = 10
var uploadMethod = "max"

type Client struct {
	w3Client              *web3go.Client
	storageUploadUrgs     *config.UploadArgs
	indexerStandardClient *indexer.Client
	indexerTurboClient    *indexer.Client
	logger                log.Logger
	MaxGasPrice           *big.Int
	NRetries              int
	Method                string
}

func New(config *config.Config, logger log.Logger) (*Client, error) {
	zgConfig, err := chain.New0gNetwork(&config.Networks)
	if err != nil {
		panic(err)
	}

	wallets, err := zgConfig.Wallets()
	if err != nil {
		panic(err)
	}
	wallet, err := wallets.Wallet(0)
	if err != nil {
		panic(err)
	}

	logger.WithFields(logrus.Fields{
		"wallet": wallet.Address(),
		"url":    zgConfig.URL(),
	}).Info("Wallet and URL")

	w3client := blockchain.MustNewWeb3(zgConfig.URL(), wallet.PrivateKey(), config.ProviderOption)
	if config.GasPrice != "" {
		gasPrice, err := strconv.ParseUint(config.GasPrice, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid gas price: %v", err)
		}
		blockchain.CustomGasPrice = gasPrice
	}
	defer w3client.Close()

	indexerStandardClient, err := indexer.NewClient(config.StorageClientConfig.IndexerStandard, indexer.IndexerClientOption{
		ProviderOption: config.ProviderOption,
		LogOption:      common.LogOption{LogLevel: logrus.InfoLevel},
	})
	if err != nil {
		return nil, err
	}

	indexerTurboClient, err := indexer.NewClient(config.StorageClientConfig.IndexerTurbo, indexer.IndexerClientOption{
		ProviderOption: config.ProviderOption,
		LogOption:      common.LogOption{LogLevel: logrus.InfoLevel},
	})
	if err != nil {
		return nil, err
	}

	maxGasPrice, err := util.ConvertToBigInt(config.MaxGasPrice)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid max gas price: %v", config.MaxGasPrice)
	}

	return &Client{
		w3Client:              w3client,
		storageUploadUrgs:     &config.StorageClientConfig.UploadArgs,
		indexerStandardClient: indexerStandardClient,
		indexerTurboClient:    indexerTurboClient,
		logger:                logger,
		MaxGasPrice:           maxGasPrice,
		NRetries:              nRetriesToUpload,
		Method:                uploadMethod,
	}, nil
}

func (c *Client) DownloadFromStorage(ctx context.Context, hash, filePath string, isTurbo bool) (string, error) {
	var indexerClient *indexer.Client
	if isTurbo {
		indexerClient = c.indexerTurboClient
	} else {
		indexerClient = c.indexerStandardClient
	}
	fileName := fmt.Sprintf("%s.zip", filePath)
	if err := os.Remove(fileName); err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	c.logger.Infof("Begin downloading and unzipping %s\n, with root: %v", fileName, hash)
	if err := indexerClient.Download(context.Background(), hash, fileName, true); err != nil {
		err = errors.Wrapf(err, "Error downloading data with root: %v", hash)
		c.logger.Errorf("%v", err)
		return "", err
	}

	topLevelDir, err := util.Unzip(fileName, filepath.Dir(filePath))
	if err != nil {
		c.logger.Errorf("Error unzipping data: %v\n", err)
		return "", err
	}

	c.logger.Infof("Downloaded and unzipped %s\n", fileName)

	return topLevelDir, nil
}

func (c *Client) UploadToStorage(ctx context.Context, fileName string, isTurbo bool) ([]ethcommon.Hash, error) {
	finalityRequired := transfer.TransactionPacked
	if c.storageUploadUrgs.FinalityRequired {
		finalityRequired = transfer.FileFinalized
	}

	opt := transfer.UploadOption{
		Tags:             hexutil.MustDecode(c.storageUploadUrgs.Tags),
		FinalityRequired: finalityRequired,
		TaskSize:         c.storageUploadUrgs.TaskSize,
		ExpectedReplica:  c.storageUploadUrgs.ExpectedReplica,
		SkipTx:           c.storageUploadUrgs.SkipTx,
		MaxGasPrice:      c.MaxGasPrice,
		NRetries:         c.NRetries,
		Method:           c.Method,
	}

	file, err := core.Open(fileName)
	if err != nil {
		c.logger.Errorf("Error opening file to upload: %v\n", err)
		return nil, err
	}
	defer file.Close()

	var indexerClient *indexer.Client
	if isTurbo {
		indexerClient = c.indexerTurboClient
	} else {
		indexerClient = c.indexerStandardClient
	}

	uploader, err := indexerClient.NewUploaderFromIndexerNodes(ctx, file.NumSegments(), c.w3Client, opt.ExpectedReplica, nil, c.Method)
	if err != nil {
		c.logger.Errorf("Error creating uploader: %v\n", err)
		return nil, err
	}
	defer indexerClient.Close()

	uploader.WithRoutines(c.storageUploadUrgs.Routines)

	_, roots, err := uploader.SplitableUpload(ctx, file, c.storageUploadUrgs.FragmentSize, opt)
	if err != nil {
		err = errors.Wrapf(err, "Error uploading file: %v", fileName)
		c.logger.Errorf("%v", err)
		return nil, err
	}
	if len(roots) == 1 {
		c.logger.Infof("file uploaded in 1 fragment, root = %v", roots[0].String())
	} else {
		s := make([]string, len(roots))
		for i, root := range roots {
			s[i] = root.String()
		}
		c.logger.Infof("file uploaded in %v fragments, roots = %v", len(roots), s)
	}

	return roots, nil
}
