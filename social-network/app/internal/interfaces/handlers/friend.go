package handlers

import (
	"net/http"

	"social_network/internal/usecases"

	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
	friendUseCase *usecases.FriendUseCase
}

func NewFriendHandler(friendUseCase *usecases.FriendUseCase) *FriendHandler {
	return &FriendHandler{friendUseCase: friendUseCase}
}

func (h *FriendHandler) AddFriend(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	friendID := c.Param("user_id")

	if err := h.friendUseCase.AddFriend(userID, friendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add friend"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *FriendHandler) DeleteFriend(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	friendID := c.Param("user_id")

	if err := h.friendUseCase.DeleteFriend(userID, friendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
