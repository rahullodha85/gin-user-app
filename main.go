package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahullodha85/gin-user-app/entity"
	"github.com/rahullodha85/gin-user-app/service"
)

func main() {
	handler := gin.Default()
	
	// healthcheck
	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"AppName": "User-App",
			"Version": "1.0.0",
		})
	})
	var videoService service.VideoServiceObj
	handler.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoService.FindAll())
	})

	handler.POST("/videos", func(ctx *gin.Context) {
		var video entity.Video
		ctx.Bind(&video)
		ctx.JSON(200, videoService.Save(video))
	})

	server := &http.Server {
		Addr: "0.0.0.0:8080", 
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Print("An error has occurred, details: %w", err)
	}
}