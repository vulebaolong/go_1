package main

import (
	"fmt"
	"go-backend/internal/common/env"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/common/response"
	"go-backend/internal/delivery"
	"go-backend/internal/handler"
	"go-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	ginEngine *gin.Engine
	env       *env.Env
}

func NewApp() *App {
	env := env.New()

	// Create a Gin router with default middleware (logger and recovery)
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(middlewares.ErrorHandler)
	ginEngine.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.Error(response.NewInternalServerErrorException())
		ctx.Abort()
	}))

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
		env:       env,
	}
}

func (a *App) Start() {
	addr := fmt.Sprintf("%s:%s", a.env.Host, a.env.Port)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	a.ginEngine.Run(addr)
}
