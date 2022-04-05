package model

import "time"

type MemberAccess struct {
	MemberId     string
	UserName     string
	Password     string
	Verification string
	JoinDate     *time.Time
}

func NewMemberAccess(id, username, password, verification string, joinDate *time.Time) *MemberAccess {
	return &MemberAccess{
		id, username, password, verification, joinDate,
	}
}
