package auth

import (
	"github.com/Humiditii/receept/internal/auth/handler"
	"github.com/Humiditii/receept/internal/auth/repository"
	"github.com/Humiditii/receept/internal/auth/service"
	"github.com/Humiditii/receept/internal/config"
	"github.com/Humiditii/receept/internal/email"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthModule(db *gorm.DB, server *gin.Engine, cfg *config.AppConfig) {
	authRepo := repository.NewAuthRepository(db)
	emailService := email.NewEmailService(cfg)
	authService := service.NewAuthService(&authRepo, emailService)
	authHandler := handler.NewAuthHandler(&authService)


	publicRoutes:= server.Group("/auth")
	{
		publicRoutes.POST("/register", authHandler.Signup)
	}

	protectedRoutes := server.Group("/auth")
	protectedRoutes.Use()
	{
		protectedRoutes.GET("/profile", func(ctx *gin.Context) {})
	}

}