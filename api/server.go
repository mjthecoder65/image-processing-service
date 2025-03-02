package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mjthecoder65/image-processing-service/config"
	db "github.com/mjthecoder65/image-processing-service/db/sqlc"
	"github.com/mjthecoder65/image-processing-service/pkg/token"
)

type Server struct {
	router  *gin.Engine
	config  *config.Config
	maker   token.Maker
	queries *db.Queries
}

func NewServer(config *config.Config, conn *pgx.Conn) (*Server, error) {
	maker, err := token.NewJWTMaker(config.JWTSecret)

	if err != nil {
		return nil, err
	}

	server := &Server{
		config:  config,
		maker:   maker,
		queries: db.New(conn),
	}

	server.SetupRoutes()

	return server, nil
}

func (server *Server) Start() error {
	return server.router.Run(server.config.ServerPort)
}

func (server *Server) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/api/v1/health", server.healthcheck)

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", server.createUser)
		auth.POST("/login", server.login)
	}

	images := router.Group("/api/v1/images")
	{
		images.POST("/", server.uploadImage)
		images.GET("/:id", server.getImage)
		images.POST("/:id/transform", server.transformImage)
		images.GET("/", server.listImages)
		images.POST("/generate", server.generateImage)
	}

	server.router = router
	return router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
