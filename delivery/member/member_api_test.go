package member

import (
	"dating_app_last/model"
	"github.com/stretchr/testify/mock"
	"time"
)

var (
	waktu = time.Now().Local()
)

var dummyMember = []model.MemberAccess{
	{
		MemberId:     "fwuefuwenufew",
		UserName:     "kocag123@gmail.com",
		Password:     "nunuec",
		Verification: "N",
		JoinDate:     &waktu,
	},
}

type memberUseCaseMock struct {
	mock.Mock
}
