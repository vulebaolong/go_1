package dependency

import (
	"go-backend/internal/delivery"
	"go-backend/internal/handler"
	"go-backend/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
)

func Injection(ginEngine *gin.Engine) {
	articleUsecase := usecase_impl.NewArticleUsecase()
	articleHandler := handler.NewArticleHandler(articleUsecase)
	articleDelivery := delivery.NewArticleDelivery(articleHandler)

	demoUsecase := usecase_impl.NewDemoUsecase()
	demoHandler := handler.NewDemoHandler(demoUsecase)
	demoDelivery := delivery.NewDemoDelivery(demoHandler)

	rootDelivery := delivery.NewRootDelivery(demoDelivery, articleDelivery)
	rootDelivery.RegisterRouter(ginEngine)
}
