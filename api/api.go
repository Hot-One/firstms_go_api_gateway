package api

import (
	_ "github.com/Hot-One/firstms_go_api_gateway/api/docs"
	"github.com/Hot-One/firstms_go_api_gateway/api/handler"
	"github.com/Hot-One/firstms_go_api_gateway/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	// "github.com/swaggo/swag/example/basic/docs"
)

func SetUpApi(r *gin.Engine, h *handler.Handler, cfg config.Config) {
	// docs.SwaggerInfo.Title = cfg.ServiceName
	// docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	r.Use(customCORSMiddleware())
	r.Use(MaxAllowed(500))

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetByIdUser)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
