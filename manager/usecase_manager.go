package manager

import "dating_app_last/usecase"

type UseCaseManager interface {
	MemberSignUpUseCase() usecase.MemberRegistration
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) MemberSignUpUseCase() usecase.MemberRegistration {
	return usecase.NewMemberRegistration(u.repoManager.MemberAccessRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager,
	}
}
