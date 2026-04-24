package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessFormat struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Doc        string `json:"doc"`
}

func Success(data any, message string, code int, ctx *gin.Context) {
	if code == 0 {
		code = http.StatusOK
	}
	if message == "" {
		message = http.StatusText(code)
	}

	reuslt := SuccessFormat{
		Status:     "success",
		StatusCode: code,
		Message:    message,
		Data:       data,
		Doc:        "",
	}
	ctx.JSON(reuslt.StatusCode, reuslt)
}
