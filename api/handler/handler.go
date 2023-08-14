package handler

import (
	"github.com/Hot-One/firstms_go_api_gateway/config"
	"github.com/Hot-One/firstms_go_api_gateway/pkg/logger"
	"github.com/Hot-One/firstms_go_api_gateway/services"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services services.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, srvc services.ServiceManagerI) *Handler {
	return &Handler{
		cfg:      cfg,
		log:      log,
		services: srvc,
	}
}
