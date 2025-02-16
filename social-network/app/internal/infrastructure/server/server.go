package server

import (
	"social_network/internal/infrastructure/cache"
	"social_network/internal/infrastructure/database"
	"social_network/internal/infrastructure/logger"
	"social_network/internal/interfaces/handlers"
	"social_network/internal/interfaces/repositories"
	"social_network/internal/middleware"
	"social_network/internal/usecases"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	logger.Init()
	// Инициализация БД и кэша
	dbRepo, err := database.NewPostgresRepository(
		"postgres://social:social@pgmaster:5432/social",
		"postgres://social:social@pgslave1:5432/social",
		"postgres://social:social@pgslave2:5432/social",
	)

	if err != nil {
		logger.ErrorLogger.Fatal("Failed to connect to database:", err)
	}

	logger.InfoLogger.Println("Successfully connected to database")

	redisCache := cache.NewRedisCache("redis:6379")

	// Репозитории
	userRepo := repositories.NewUserRepository(dbRepo)
	postRepo := repositories.NewPostRepository(dbRepo)
	friendRepo := repositories.NewFriendRepository(dbRepo)

	// Use Cases
	authUseCase := usecases.NewAuthUseCase(userRepo)
	userUseCase := usecases.NewUserUseCase(userRepo)
	postUseCase := usecases.NewPostUseCase(postRepo, redisCache)
	friendUseCase := usecases.NewFriendUseCase(friendRepo)

	// Обработчики
	authHandler := handlers.NewAuthHandler(authUseCase)
	userHandler := handlers.NewUserHandler(userUseCase)
	postHandler := handlers.NewPostHandler(postUseCase)
	friendHandler := handlers.NewFriendHandler(friendUseCase)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware()

	// Настройка маршрутов
	r := gin.Default()

	// Public routes
	r.POST("/login", authHandler.Login)
	r.POST("/user/register", userHandler.RegisterUser)

	// Protected routes
	authGroup := r.Group("/")
	authGroup.Use(authMiddleware.VerifyToken())
	{
		authGroup.GET("/user/get/:id", userHandler.GetUser)
		authGroup.GET("/user/search", userHandler.SearchUsers)
		authGroup.PUT("/friend/set/:user_id", friendHandler.AddFriend)
		authGroup.PUT("/friend/delete/:user_id", friendHandler.DeleteFriend)
		authGroup.POST("/post/create", postHandler.CreatePost)
		authGroup.PUT("/post/update", postHandler.UpdatePost)
		authGroup.PUT("/post/delete/:id", postHandler.DeletePost)
		authGroup.GET("/post/get/:id", postHandler.GetPost)
		authGroup.GET("/post/feed", postHandler.GetFeed)
	}

	logger.InfoLogger.Println("Server starting on port 80")
	r.Run(":80")
}
