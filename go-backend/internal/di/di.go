package dependency

import (
	"go-backend/ent"
	"go-backend/internal/common/env"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/common/socket"
	storage_impl "go-backend/internal/common/storage/impl"
	"go-backend/internal/delivery"
	"go-backend/internal/handler"
	"go-backend/internal/repository/repository_impl"
	"go-backend/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, gormClient *gorm.DB, env *env.Env, allowOrigins []string) {
	tokenUsecase := usecase_impl.NewTokenUsecase(env)

	// repository
	articleRepository := repository_impl.NewArticleRepository(entClient, gormClient)
	userRepository := repository_impl.NewUserRepository(entClient)
	chatGroupRepository := repository_impl.NewChatGroupRepository(entClient)
	chatGroupMemberRepository := repository_impl.NewChatGroupMemberRepository(entClient)
	unitOfWorkRepository := repository_impl.NewUnitOfWorkRepository(entClient)

	localFileStorage := storage_impl.NewLocalFileStorage("public")
	cloudinaryFileStorage := storage_impl.NewCloudinaryStorage(env)

	// Socket
	chatUsecase := usecase_impl.NewChatUsecase(tokenUsecase, userRepository, chatGroupRepository, chatGroupMemberRepository, unitOfWorkRepository)
	chatHandler := handler.NewChatHandler(chatUsecase)
	socket := socket.NewSocket(chatHandler)
	socket.Start(ginEngine, allowOrigins)

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

	userUsecase := usecase_impl.NewUserUsecase(userRepository, localFileStorage, cloudinaryFileStorage)
	userHandler := handler.NewUserHandler(userUsecase)
	userDelivery := delivery.NewUserDelivery(userHandler, authMiddleware)

	chatGroupUsecase := usecase_impl.NewChatGroupUsecase(chatGroupRepository)
	chatGroupHandler := handler.NewChatGroupHandler(chatGroupUsecase)
	chatGroupDelivery := delivery.NewChatGroupDelivery(chatGroupHandler)

	rootDelivery := delivery.NewRootDelivery(demoDelivery, articleDelivery, authDelivery, userDelivery, chatGroupDelivery)
	rootDelivery.RegisterRouter(ginEngine)
}
