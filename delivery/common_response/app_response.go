package common_response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppHttpResponse interface {
	SendData(message *ResponseMessage)
	SendError(message *ErrorMessage) error
}

type JsonResponse struct {
	gx *gin.Context
}

func (j *JsonResponse) SendData(message *ResponseMessage) {
	j.gx.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendError(message *ErrorMessage) error {
	return nil
}
