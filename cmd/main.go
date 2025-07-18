package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Humiditii/receept/internal/auth"
	"github.com/Humiditii/receept/internal/auth/model"
	"github.com/Humiditii/receept/internal/config"
	"github.com/Humiditii/receept/internal/database"
	"github.com/gin-gonic/gin"
)

func main(){

	fileLogging()

	config := config.NewConfigInstance(".env")

	dbInstance := database.NewDatabase(config, &model.User{})

	defer dbInstance.CloseDBInstance()

	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"welcome to receept",
		})
	})

	auth.RegisterAuthModule(dbInstance.Conn, server)

	server.Run(":"+config.GetPort())
}


func fileLogging(){
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}