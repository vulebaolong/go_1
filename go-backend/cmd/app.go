package main

import (
	"fmt"
	"go-backend/ent"
	"go-backend/internal/common/ent_client"
	"go-backend/internal/common/env"
	"go-backend/internal/common/gorm_client"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/common/response"
	dependency "go-backend/internal/di"

	"github.com/gin-contrib/cors"
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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://google.com"}
	ginEngine.Use(cors.New(corsConfig))

	// ginEngine.Use(func(ctx *gin.Context) {
	// 	ctx.Header("access-control-allow-origin", "http://localhost:3000")
	// 	if ctx.Request.Method == "OPTIONS" {
	// 		ctx.Header("access-control-allow-headers", "Origin,Content-Length,Content-Type")
	// 		ctx.Header("access-control-allow-methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
	// 		ctx.Header("access-control-max-age", "43201")
	// 		ctx.AbortWithStatus(204)
	// 		return
	// 	}
	// })

	// ginEngine.Use(middlewares.A)
	// ginEngine.Use(middlewares.B)
	// ginEngine.Use(middlewares.C)
	entClient := ent_client.New(env)
	gormClient := gorm_client.New(env)
	dependency.Injection(ginEngine, entClient, gormClient)

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
