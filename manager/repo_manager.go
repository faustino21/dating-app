package manager

import "dating_app_last/repository"

type RepoManager interface {
	MemberAccessRepo() repository.MemberAccessRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) MemberAccessRepo() repository.MemberAccessRepo {
	return repository.NewMemberAccessRepo(r.infra.GetSqlConn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
