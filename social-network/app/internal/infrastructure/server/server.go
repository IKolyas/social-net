package server

import (
	"os"
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

	// "postgres://social:social@pgmaster:5432/social"
	// "postgres://social:social@pgslave1:5432/social"
	// "postgres://social:social@pgslave2:5432/social"
	DB_MASTER := os.Getenv("DB_MASTER")
	DB_SLAVE1 := os.Getenv("DB_SLAVE1")
	DB_SLAVE2 := os.Getenv("DB_SLAVE2")
	REDIS_CONNECTION := os.Getenv("REDIS_CONNECTION")
	SERVER_PORT := os.Getenv("SERVER_PORT")

	// Инициализация БД и кэша
	dbRepo, err := database.NewPostgresRepository(DB_MASTER, DB_SLAVE1, DB_SLAVE2)

	if err != nil {
		logger.ErrorLogger.Printf("Failed to connect to database: %s", err)
	}

	logger.InfoLogger.Println("Successfully connected to database")

	//"redis:6379"
	redisCache := cache.NewRedisCache(REDIS_CONNECTION)

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
	r.GET("/health", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": "ok"}) })
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
	r.Run(":" + SERVER_PORT)
}
