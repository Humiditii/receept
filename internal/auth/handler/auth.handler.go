package handler

import (
	"net/http"

	"github.com/Humiditii/receept/internal/auth/model"
	"github.com/Humiditii/receept/internal/auth/service"
	"github.com/Humiditii/receept/pkg/utils"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService *service.AuthService
}

type AuthHandler interface {
	Signup(ctx *gin.Context)
}


func NewAuthHandler (service *service.AuthService) AuthHandler {
	return & authHandler{
		authService: service,
	}
}

func (a *authHandler) Signup (ctx *gin.Context) {


	var user model.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		
		utils.Send(ctx, http.StatusBadRequest, err.Error(), nil, true)

		return
	}

	result, err := (*a.authService).Signup(&user)

	if err != nil {
		utils.Send(ctx, http.StatusBadRequest, err.Error(), nil, true)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"user created!",
		"data": result,
	})

	utils.Send(ctx, http.StatusCreated, "New user created", nil, false)
}