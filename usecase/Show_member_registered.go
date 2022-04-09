package usecase

import (
	"dating_app_last/model"
	"dating_app_last/repository"
)

type ShowMemberRegistered interface {
	GetAllMember(page int) (*[]model.GetMemberResp, error)
}

type showMemberRegistered struct {
	repo repository.MemberAccessRepo
}

func (s *showMemberRegistered) GetAllMember(page int) (*[]model.GetMemberResp, error) {
	return s.repo.GetMember(page)
}

func NewShowMemberRegistered(repo repository.MemberAccessRepo) ShowMemberRegistered {
	return &showMemberRegistered{
		repo: repo,
	}
}
