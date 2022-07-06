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

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	// Get Posts from open api
	api.GET("/get-from-open-api/", handlerV1.GetPostsFromApi)

	// Posts
	api.POST("/posts/create-post/", handlerV1.CreatePost)
	api.GET("/posts/get-post/:id/", handlerV1.GetPostById)
	api.PUT("/posts/update-post/:id/", handlerV1.UpdatePost)
	api.DELETE("/posts/delete-post/:id/", handlerV1.DeletePost)
	api.GET("/posts/list-posts/", handlerV1.ListPosts)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
