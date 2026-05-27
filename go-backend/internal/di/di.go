package dependency

import (
	"go-backend/ent"
	"go-backend/internal/common/env"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/delivery"
	"go-backend/internal/handler"
	"go-backend/internal/repository/repository_impl"
	"go-backend/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, gormClient *gorm.DB, env *env.Env) {
	articleRepository := repository_impl.NewArticleRepository(entClient, gormClient)
	userRepository := repository_impl.NewUserRepository(entClient)

	tokenUsecase := usecase_impl.NewTokenUsecase(env)
	authMiddleware := middlewares.NewAuthMiddleware(tokenUsecase, userRepository)

	articleUsecase := usecase_impl.NewArticleUsecase(articleRepository)
	articleHandler := handler.NewArticleHandler(articleUsecase)
	articleDelivery := delivery.NewArticleDelivery(articleHandler)

	demoUsecase := usecase_impl.NewDemoUsecase()
	demoHandler := handler.NewDemoHandler(demoUsecase)
	demoDelivery := delivery.NewDemoDelivery(demoHandler)

	authUsecase := usecase_impl.NewAuthUsecase(userRepository, tokenUsecase, env)
	authHandler := handler.NewAuthHandler(authUsecase, env)
	authDelivery := delivery.NewAuthDelivery(authHandler, authMiddleware)

	userUsecase := usecase_impl.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	userDelivery := delivery.NewUserDelivery(userHandler, authMiddleware)

	rootDelivery := delivery.NewRootDelivery(demoDelivery, articleDelivery, authDelivery, userDelivery)
	rootDelivery.RegisterRouter(ginEngine)
}
