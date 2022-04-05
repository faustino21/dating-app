package usecase

import (
	"dating_app_last/repository"
	"time"
)

type MemberRegistration interface {
	MemberSignUp(username, password, memberId string, joinDate *time.Time, verification string) error
}

type memberRegistration struct {
	repo repository.MemberAccessRepo
}

func (m *memberRegistration) MemberSignUp(username, password, memberId string, joinDate *time.Time, verification string) error {
	err := m.repo.Insert(username, password, memberId, joinDate, verification)
	if err != nil {
		return err
	}
	return nil
}

func NewMemberRegistration(repo repository.MemberAccessRepo) MemberRegistration {
	return &memberRegistration{
		repo: repo,
	}
}
