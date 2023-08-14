package main

import (
	"fmt"

	"github.com/Hot-One/firstms_go_api_gateway/api"
	"github.com/Hot-One/firstms_go_api_gateway/api/handler"
	"github.com/Hot-One/firstms_go_api_gateway/config"
	"github.com/Hot-One/firstms_go_api_gateway/pkg/logger"
	"github.com/Hot-One/firstms_go_api_gateway/services"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	fmt.Printf("config: %+v/n", cfg)

	// Setup Logger
	loggerLevel := logger.LevelDebug
	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	grpcSrvc, err := services.NewgrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	h := handler.NewHandler(cfg, log, grpcSrvc)

	api.SetUpApi(r, h, cfg)

	fmt.Println("Start api gateway...")

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		return
	}
}
