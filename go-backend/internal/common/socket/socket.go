package socket

import (
	"fmt"
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
	server "github.com/zishang520/socket.io/servers/socket/v3"
	"github.com/zishang520/socket.io/v3/pkg/types"
)

type socket struct {
	chatHandler *handler.ChatHandler
}

func NewSocket(chatHandler *handler.ChatHandler) *socket {
	return &socket{
		chatHandler: chatHandler,
	}
}

// https://github.com/zishang520/socket.io/blob/v3/docs/UPGRADE.md#quick-start-example
func (s *socket) Start(ginEngine *gin.Engine, allowOrigins []string) {
	option := server.DefaultServerOptions()

	option.SetCors(&types.Cors{
		Origin: allowOrigins,
	})

	io := server.NewServer(nil, option)

	io.On("connection", func(args ...any) {
		socket := args[0].(*server.Socket)
		fmt.Printf("connected: %s\n", socket.Id())

		socket.On("CREATE_ROOM", func(args ...any) {
			// handler
			s.chatHandler.CreateGroup()
			fmt.Printf("received: %v\n", args)
			socket.Emit("message", args...)
		})

		socket.On("disconnect", func(args ...any) {
			fmt.Printf("disconnected: %s\n", socket.Id())
		})
	})

	ginEngine.Any("/socket.io/", gin.WrapH(io.ServeHandler(nil)))
}
