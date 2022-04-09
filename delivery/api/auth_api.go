package api

import (
	"dating_app_last/delivery/commonResponse"
	"dating_app_last/delivery/httpRequest"
	"dating_app_last/delivery/middleware"
	"dating_app_last/model"
	"dating_app_last/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthApi struct {
	commonResponse.BaseApi
	memberAuth usecase.Authentication
}

func (a *AuthApi) AuthenticationMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var memberReq httpRequest.MemberRequest
		var member *model.MemberAccess
		err := a.ParseRequestBody(c, &memberReq)
		if err != nil {
			a.ParsingError(c, err)
			return
		}
		member, err = a.memberAuth.Authentication(memberReq)
		if err != nil {
			commonResponse.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResponse.NewFailedMessage(err.Error()))
			return
		}
		commonResponse.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResponse.NewSuccessMessage("Authentication Success", member))
		token, err := middleware.GenerateToken(member.UserName, member.JoinDate)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	}
}

func NewAuthentication(memberRoute *gin.RouterGroup, authentication usecase.Authentication) (*AuthApi, error) {
	authApi := AuthApi{
		memberAuth: authentication,
	}
	memberRoute.POST("/login", authApi.AuthenticationMember())
	return &authApi, nil
}
