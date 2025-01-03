package handler

import (
	"github.com/0glabs/0g-serving-broker/common/errors"
	"github.com/0glabs/0g-serving-broker/fine-tuning/config"
	"github.com/0glabs/0g-serving-broker/fine-tuning/internal/ctrl"
	"github.com/0glabs/0g-storage-client/common"
	"github.com/0glabs/0g-storage-client/indexer"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	ctrl *ctrl.Ctrl
	indexerStandardClient *indexer.Client
	indexerTurboClient *indexer.Client
}

func New(ctrl *ctrl.Ctrl, config *config.Config) *Handler {
	indexerStandardClient, err := indexer.NewClient(config.IndexerStandardUrl, indexer.IndexerClientOption{
		ProviderOption: config.ProviderOption,
		LogOption:      common.LogOption{Logger: logrus.StandardLogger()},
	})
	if err != nil {
		return nil
	}

	indexerTurboClient, err := indexer.NewClient(config.IndexerTurboUrl, indexer.IndexerClientOption{
		ProviderOption: config.ProviderOption,
		LogOption:      common.LogOption{Logger: logrus.StandardLogger()},
	})
	if err != nil {
		return nil
	}

	h := &Handler{
		ctrl: ctrl,
		indexerStandardClient: indexerStandardClient,
		indexerTurboClient: indexerTurboClient,
	}
	return h
}

func (h *Handler) Register(r *gin.Engine) {
	group := r.Group("/v1")

	group.POST("/task", h.CreateTask)
	group.GET("/task/:taskID", h.GetTask)
}

func handleBrokerError(ctx *gin.Context, err error, context string) {
	info := "Provider"
	if context != "" {
		info += (": " + context)
	}
	errors.Response(ctx, errors.Wrap(err, info))
}
