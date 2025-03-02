package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mjthecoder65/image-processing-service/config"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer(config *config.Config) (*Server, error) {
	router := gin.Default()

	router.POST("/api/auth/register", func(ctx *gin.Context) {

	})

	router.POST("/api/auth/login", func(ctx *gin.Context) {

	})

	router.POST("/api/v1/images", func(ctx *gin.Context) {

	})

	router.GET("/api/v1/images/:id", func(ctx *gin.Context) {

	})

	router.POST("/api/v1/images/:id/transform", func(ctx *gin.Context) {

	})

	router.GET("/api/v1/images", func(ctx *gin.Context) {

	})

	return &Server{
		config: config,
		router: router,
	}, nil
}

func (server *Server) Start() error {
	return server.router.Run(server.config.ServerPort)
}
