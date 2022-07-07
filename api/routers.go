package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Muhammadjon226/api_gateway/api/docs" // swag
	v1 "github.com/Muhammadjon226/api_gateway/api/handlers/v1"
	"github.com/Muhammadjon226/api_gateway/config"
	"github.com/Muhammadjon226/api_gateway/pkg/logger"
	"github.com/Muhammadjon226/api_gateway/services"
)

// Config ...
type Config struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(cfg Config) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	HandlerV1 := v1.New(
		cfg.Logger,
		cfg.ServiceManager,
		cfg.Conf)

	api := router.Group("/v1")
	// Get Posts from open api
	api.GET("/posts/get-from-open-api/", HandlerV1.GetPostsFromAPI)

	// Posts
	api.POST("/posts/create-post/", HandlerV1.CreatePost)
	api.GET("/posts/get-post/:id/", HandlerV1.GetPostByID)
	api.PUT("/posts/update-post/:id/", HandlerV1.UpdatePost)
	api.DELETE("/posts/delete-post/:id/", HandlerV1.DeletePost)
	api.GET("/posts/list-posts/", HandlerV1.ListPosts)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
