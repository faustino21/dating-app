package model

import "time"

type GetMemberResp struct {
	UserName string     `db:"user_name"`
	JoinDate *time.Time `db:"join_date"`
}

func NewGetMemberResp(username string, joinDate *time.Time) GetMemberResp {
	return GetMemberResp{
		UserName: username,
		JoinDate: joinDate,
	}
}
