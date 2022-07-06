package v1

import (
	"time"
	"context"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Muhammadjon226/api_gateway/api/handlers/models"
	
	pbPost "github.com/Muhammadjon226/api_gateway/genproto/post_service"
	l "github.com/Muhammadjon226/api_gateway/pkg/logger"
	"github.com/Muhammadjon226/api_gateway/pkg/utils"
)

// CreatePost ...
// @Summary CreatePost
// @Description This API for creating a new post
// @Tags post
// @Accept  json
// @Produce  json
// @Param create_post_request body models.Post true "create_post_request"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/create-post/ [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body models.Post
	)
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().CreatePost(ctx, &pbPost.Post{
		Id:     body.Id,
		UserId: body.UserId,
		Title:  body.Title,
		Body:   body.Body,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetPostById ...
// @Summary GetPostById
// @Description This API for getting post by id
// @Tags post
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/get-post/{id}/ [get]
func (h *handlerV1) GetPostById(c *gin.Context) {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().GetPostById(ctx, &pbPost.ByIdReq{
		Id: int64(postId),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post by id", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdatePost ...
// @Summary UpdatePost
// @Description This API for updating post
// @Tags post
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param update_post body models.UpdatePost true "update_post"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/update-post/{id}/ [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body models.UpdatePost
	)
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id", l.Error(err))
		return
	}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().CreatePost(ctx, &pbPost.Post{
		Id:     int64(postId),
		UserId: body.UserId,
		Title:  body.Title,
		Body:   body.Body,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update post", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeletePost ...
// @Summary DeletePost
// @Description This API for deleting post
// @Tags post
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.SuccessfullResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/delete-post/{id}/ [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	_, err = h.serviceManager.PostService().DeletePost(ctx, &pbPost.ByIdReq{
		Id: int64(postId),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, models.SuccessfullResponse{
		Message: "Post successfully deleted",
	})
}

// ListPosts ...
// @Summary ListPosts
// @Description This API for getting list of posts
// @Tags post
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListPosts
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/list-posts/ [get]
func (h *handlerV1) ListPosts(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().ListPosts(ctx, &pbPost.ListReq{
		Page:  params.Page,
		Limit: params.Limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list posts", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}