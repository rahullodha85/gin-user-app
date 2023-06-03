package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
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

	server := &http.Server {
		Addr: "0.0.0.0:8080", 
		Handler: handler,
	}


	err := server.ListenAndServe()
	if err != nil {
		fmt.Print("An error has occurred, details: %w", err)
	}
}