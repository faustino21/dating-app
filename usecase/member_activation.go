package usecase

import "dating_app_last/repository"

type MemberActivationUseCase interface {
	MemberActivation(id string) error
}

type memberActivationUseCase struct {
	repo repository.MemberAccessRepo
}

func (m *memberActivationUseCase) MemberActivation(id string) error {
	return m.repo.Update(id)
}

func NewActivationUseCase(repo repository.MemberAccessRepo) MemberActivationUseCase {
	return &memberActivationUseCase{
		repo: repo,
	}
}
