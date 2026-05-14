package dependency

import (
	"go-backend/ent"
	"go-backend/internal/delivery"
	"go-backend/internal/handler"
	"go-backend/internal/repository/repository_impl"
	"go-backend/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, gormClient *gorm.DB) {
	articleRepository := repository_impl.NewArticleRepository(entClient, gormClient)
	userRepository := repository_impl.NewUserRepository(entClient)

	articleUsecase := usecase_impl.NewArticleUsecase(articleRepository)
	articleHandler := handler.NewArticleHandler(articleUsecase)
	articleDelivery := delivery.NewArticleDelivery(articleHandler)

	demoUsecase := usecase_impl.NewDemoUsecase()
	demoHandler := handler.NewDemoHandler(demoUsecase)
	demoDelivery := delivery.NewDemoDelivery(demoHandler)

	authUsecase := usecase_impl.NewAuthUsecase(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)
	authDelivery := delivery.NewAuthDelivery(authHandler)

	rootDelivery := delivery.NewRootDelivery(demoDelivery, articleDelivery, authDelivery)
	rootDelivery.RegisterRouter(ginEngine)
}
