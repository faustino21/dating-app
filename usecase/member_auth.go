package usecase

import (
	"dating_app_last/delivery/httpRequest"
	"dating_app_last/model"
	"dating_app_last/repository"
)

type Authentication interface {
	Authentication(memberReq httpRequest.MemberRequest) (*model.MemberAccess, error)
}

type authentication struct {
	repo repository.MemberAccessRepo
}

func (a *authentication) Authentication(memberReq httpRequest.MemberRequest) (*model.MemberAccess, error) {
	return a.repo.Login(memberReq.Email, memberReq.Password)
}

func NewAuthentication(repo repository.MemberAccessRepo) Authentication {
	return &authentication{
		repo: repo,
	}
}
