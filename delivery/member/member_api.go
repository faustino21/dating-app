package member

import (
	"dating_app_last/delivery/common_response"
	"dating_app_last/usecase"
	"dating_app_last/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MemberApi struct {
	common_response.BaseApi
	memberSignUp usecase.MemberRegistration
}

func (m *MemberApi) SignUpMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var memberReq MemberRequest
		err := m.ParseRequestBody(c, &memberReq)
		if err != nil {
			util.Log.Error().Msg("Parse Error")
			return
		}
		timeStamp := time.Now().Local()
		uuid := util.GetUuid()
		err = m.memberSignUp.MemberSignUp(memberReq.Email, memberReq.Password, uuid, &timeStamp, "N")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		}
		c.String(http.StatusOK, fmt.Sprintf("Success"))

		c.
	}
}



func NewMemberApi(memberRoute *gin.RouterGroup, memberSignUp usecase.MemberRegistration) (*MemberApi, error) {
	memberApi := MemberApi{
		memberSignUp: memberSignUp,
	}

	memberGroup := memberRoute.Group("/member")
	memberGroup.POST("/signup", memberApi.SignUpMember())
	return &memberApi, nil
}
