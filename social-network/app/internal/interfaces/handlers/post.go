package handlers

import (
	"net/http"
	"strconv"

	"social_network/internal/entities"
	"social_network/internal/usecases"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUseCase *usecases.PostUseCase
}

func NewPostHandler(postUseCase *usecases.PostUseCase) *PostHandler {
	return &PostHandler{postUseCase: postUseCase}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post entities.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	post.AuthorUserID = c.MustGet("user_id").(string)
	if err := h.postUseCase.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post_id": post.ID})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	var post entities.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.postUseCase.UpdatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postID := c.Param("id")
	if err := h.postUseCase.DeletePost(postID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *PostHandler) GetPost(c *gin.Context) {
	postID := c.Param("id")
	post, err := h.postUseCase.GetPost(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) GetFeed(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	userID := c.MustGet("user_id").(string)
	posts, err := h.postUseCase.GetFeed(userID, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get feed"})
		return
	}

	c.JSON(http.StatusOK, posts)

}
