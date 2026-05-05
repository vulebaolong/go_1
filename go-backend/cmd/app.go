package main

import (
	"fmt"
	"go-backend/ent"
	"go-backend/internal/common/ent_client"
	"go-backend/internal/common/env"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/common/response"
	dependency "go-backend/internal/di"

	"github.com/gin-gonic/gin"
)

type App struct {
	ginEngine *gin.Engine
	env       *env.Env
	entClient *ent.Client
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

	// ginEngine.Use(middlewares.A)
	// ginEngine.Use(middlewares.B)
	// ginEngine.Use(middlewares.C)
	entClient := ent_client.New()
	dependency.Injection(ginEngine)

	return &App{
		ginEngine: ginEngine,
		env:       env,
		entClient: entClient,
	}
}

func (a *App) Start() {
	addr := fmt.Sprintf("%s:%s", a.env.Host, a.env.Port)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	a.ginEngine.Run(addr)
}
