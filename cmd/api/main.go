package main

import (
    "log"

    "github.com/username/go-base-project/internal/config"
    "github.com/username/go-base-project/internal/controller"
    "github.com/username/go-base-project/internal/handler"
    "github.com/username/go-base-project/internal/middleware"
    "github.com/username/go-base-project/internal/model"
    "github.com/username/go-base-project/internal/repository"
    "github.com/username/go-base-project/internal/service"
    "github.com/username/go-base-project/pkg/database"
    "github.com/username/go-base-project/pkg/logger"

    "github.com/gin-gonic/gin"
)

func main() {
    // 1. Load Config
    cfg := config.LoadConfig()

    // 2. Setup Logger
    logger.InitLogger(cfg.AppEnv)

    // 3. Connect to Database
    db, err := database.InitDB(cfg)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // 4. Run Migrations
    if err := db.AutoMigrate(&model.User{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    // 5. Dependency Injection
    userRepo := repository.NewUserRepository(db)
    authService := service.NewAuthService(userRepo, cfg)
    authController := controller.NewAuthController(authService)
    authHandler := handler.NewAuthHandler(authController)
    pingHandler := handler.NewPingHandler()

    // 6. Setup Router
    r := gin.Default()
    r.Use(logger.LoggerMiddleware())

    v1 := r.Group("/api/v1")
    {
        // Public Routes
        v1.GET("/ping", pingHandler.Ping)
        authGroup := v1.Group("/auth")
        {
            authGroup.POST("/register", authHandler.Register)
            authGroup.POST("/login", authHandler.Login)
        }

        // Protected Routes
        protected := v1.Group("/")
        protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
        {
            protected.GET("/profile", authHandler.Profile)
        }
    }

    // 7. Start Server
    port := ":" + cfg.AppPort
    log.Printf("Server starting on port %s", port)
    if err := r.Run(port); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}
