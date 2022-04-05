package commonResponse

import (
	"github.com/gin-gonic/gin"
)

type AppHttpResponse interface {
	SuccessResp(httpCode int, successMessage *ResponseMessage)
	FailedResp(httpCode int, failedMessage *FailedMessage)
}

type JsonResponse struct {
	gx *gin.Context
}

func (j *JsonResponse) SuccessResp(httpCode int, successMessage *ResponseMessage) {
	successMessage.HttpResponse = httpCode
	j.gx.JSON(httpCode, successMessage)
	j.gx.Abort()
}

func (j *JsonResponse) FailedResp(httpCode int, failedMessage *FailedMessage) {
	failedMessage.HttpResponse = httpCode
	j.gx.JSON(httpCode, failedMessage)
	j.gx.Abort()
}

func NewAppHttpResponse(gx *gin.Context) AppHttpResponse {
	return &JsonResponse{
		gx,
	}
}

//	j.gx.JSON(http.StatusBadRequest, gin.H{
//		"message" : "hello",
//	})
