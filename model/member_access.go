package model

import "time"

type MemberAccess struct {
	MemberId     string
	UserName     string
	Password     string
	Verification string
	JoinDate     *time.Time
}
