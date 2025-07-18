package auth

import (
	"github.com/Humiditii/receept/internal/auth/handler"
	"github.com/Humiditii/receept/internal/auth/repository"
	"github.com/Humiditii/receept/internal/auth/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthModule(db *gorm.DB, server *gin.Engine) {
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(&authRepo)
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