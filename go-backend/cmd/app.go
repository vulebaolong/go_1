package main

import (
	"go-backend/internal/delivery"
	"go-backend/internal/handler"
	"go-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	ginEngine *gin.Engine
}

func NewApp() *App {

	// Create a Gin router with default middleware (logger and recovery)
	ginEngine := gin.Default()

	demoUsecase := usecase.NewDemoUsecase()
	demoHandler := handler.NewDemoHandler(demoUsecase)
	demoDelivery := delivery.NewDemoDelivery(demoHandler)

	rootDelivery := delivery.NewRootDelivery(demoDelivery)
	rootDelivery.RegisterRouter(ginEngine)

	// Define a simple GET endpoint
	// rest API, restful API
	ginEngine.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message1": "pong",
		})
	})

	return &App{
		ginEngine: ginEngine,
	}
}

func (a *App) Start() {
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	a.ginEngine.Run()
}
