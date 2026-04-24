package handler

import (
	"go-backend/internal/common/response"
	"go-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	demoUsecase *usecase.DemoUsecase
}

func NewDemoHandler(demoUsecase *usecase.DemoUsecase) *DemoHandler {
	return &DemoHandler{
		demoUsecase: demoUsecase,
	}
}

func (d *DemoHandler) Query(ctx *gin.Context) {
	// tham số nhận vào
	data := d.demoUsecase.Query()
	response.Success(data, "", 0, ctx)
}
