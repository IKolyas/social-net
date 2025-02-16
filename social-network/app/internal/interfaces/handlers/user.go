package handlers

import (
	"net/http"

	"social_network/internal/entities"
	"social_network/internal/infrastructure/logger"
	"social_network/internal/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

type RegisterRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Birthdate  string `json:"birthdate"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	Password   string `json:"password"`
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.ErrorLogger.Printf("Invalid request data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user := entities.User{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Birthdate:  req.Birthdate,
		Biography:  req.Biography,
		City:       req.City,
		Password:   req.Password,
	}

	if err := h.userUseCase.Register(&user); err != nil {
		logger.ErrorLogger.Printf("Failed to register user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	logger.InfoLogger.Printf("User registered successfully: %s", user.ID)
	c.JSON(http.StatusOK, gin.H{"user_id": user.ID})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.userUseCase.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	query := c.Query("search")

	users, err := h.userUseCase.SearchUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
