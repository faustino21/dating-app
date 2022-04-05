package manager

import "dating_app_last/usecase"

type UseCaseManager interface {
	MemberSignUpUseCase() usecase.MemberRegistration
	MemberActivationUseCase() usecase.MemberActivationUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) MemberSignUpUseCase() usecase.MemberRegistration {
	return usecase.NewMemberRegistration(u.repoManager.MemberAccessRepo())
}

func (u *useCaseManager) MemberActivationUseCase() usecase.MemberActivationUseCase {
	return usecase.NewActivationUseCase(u.repoManager.MemberAccessRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager,
	}
}
