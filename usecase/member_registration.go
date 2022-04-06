package usecase

import (
	"dating_app_last/model"
	"dating_app_last/repository"
	"time"
)

type MemberRegistration interface {
	MemberSignUp(username, password, memberId string, joinDate *time.Time, verification string) (model.MemberAccess, error)
}

type memberRegistration struct {
	repo repository.MemberAccessRepo
}

func (m *memberRegistration) MemberSignUp(username, password, memberId string, joinDate *time.Time, verification string) (model.MemberAccess, error) {
	newMember := model.NewMemberAccess(memberId, username, password, verification, joinDate)
	err := m.repo.Insert(newMember.UserName, newMember.Password, newMember.MemberId, newMember.JoinDate, newMember.Verification)
	if err != nil {
		return *newMember, err
	}
	return *newMember, nil
}

func NewMemberRegistration(repo repository.MemberAccessRepo) MemberRegistration {
	return &memberRegistration{
		repo: repo,
	}
}
