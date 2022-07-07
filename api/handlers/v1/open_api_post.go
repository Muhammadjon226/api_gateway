package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/Muhammadjon226/api_gateway/api/handlers/models"
	pbFirst "github.com/Muhammadjon226/api_gateway/genproto/first_service"
	l "github.com/Muhammadjon226/api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

// GetPostsFromAPI ...
// @Summary GetPostsFromOpenAPI
// @Description This API for getting list of posts from open api
// @Tags open_api_posts
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SuccessfullResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/get-from-open-api/ [get]
func (h *HandlerV1) GetPostsFromAPI(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	_, err := h.serviceManager.FirstService().GetPostsFromOpenAPI(ctx, &pbFirst.EmptyResp{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list posts from open api", l.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, models.SuccessfullResponse{
		Message: "Post written to database successfully from open api",
	})	
}
