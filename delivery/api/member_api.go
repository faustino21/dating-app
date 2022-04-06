package api

import (
	"dating_app_last/delivery/commonResponse"
	"dating_app_last/delivery/httpRequest"
	"dating_app_last/model"
	"dating_app_last/usecase"
	"dating_app_last/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MemberApi struct {
	commonResponse.BaseApi
	memberSignUp     usecase.MemberRegistration
	memberActivation usecase.MemberActivationUseCase
}

func (m *MemberApi) SignUpMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var memberReq httpRequest.MemberRequest
		var value model.MemberAccess
		err := m.ParseRequestBody(c, &memberReq)
		if err != nil {
			m.ParsingError(c, err)
			return
		}
		timeStamp := time.Now().Local()
		uuid := util.GetUuid()
		value, err = m.memberSignUp.MemberSignUp(memberReq.Email, memberReq.Password, uuid, &timeStamp, "N")
		if err != nil {
			commonResponse.NewAppHttpResponse(c).FailedResp(http.StatusInternalServerError, commonResponse.NewFailedMessage(err.Error()))
			return

		}
		commonResponse.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResponse.NewSuccessMessage("Sign Up Success", value))
	}
}

func (m *MemberApi) ActivationMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var memberId httpRequest.MemberIdReq
		err := m.ParseRequestBody(c, &memberId)
		if err != nil {
			m.ParsingError(c, err)
			return
		}
		err = m.memberActivation.MemberActivation(memberId.MemberId)
		if err != nil {
			commonResponse.NewAppHttpResponse(c).FailedResp(http.StatusInternalServerError, commonResponse.NewFailedMessage(err.Error()))
			return
		}
		commonResponse.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResponse.NewSuccessMessage("Success", "Verified"))
	}

}

func NewMemberApi(memberRoute *gin.RouterGroup, memberSignUp usecase.MemberRegistration, memberActivation usecase.MemberActivationUseCase) (*MemberApi, error) {
	memberApi := MemberApi{
		memberSignUp:     memberSignUp,
		memberActivation: memberActivation,
	}

	memberGroup := memberRoute.Group("/api")
	memberGroup.POST("/signup", memberApi.SignUpMember())
	memberGroup.POST("/verification", memberApi.ActivationMember())
	return &memberApi, nil
}
