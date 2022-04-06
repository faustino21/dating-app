package model

import "time"

type MemberAccess struct {
	MemberId     string     `json:"member_id" db:"member_id"`
	UserName     string     `json:"user_name" db:"user_name"`
	Password     string     `json:"password" db:"user_password"`
	Verification string     `json:"verification" db:"join_date"`
	JoinDate     *time.Time `json:"join_date" db:"verification_status"`
}

func NewMemberAccess(id, username, password, verification string, joinDate *time.Time) *MemberAccess {
	return &MemberAccess{
		MemberId:     id,
		UserName:     username,
		Password:     password,
		Verification: verification,
		JoinDate:     joinDate,
	}
}
