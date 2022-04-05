package repository

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type MemberAccessRepo interface {
	Insert(username, password, memberId string, joinDate *time.Time, verification string) error
}

type memberAccessRepoImpl struct {
	db *sqlx.DB
}

func (m *memberAccessRepoImpl) Insert(username, password, memberId string, joinDate *time.Time, verification string) error {
	_, err := m.db.Exec("insert into member_access(user_name, user_password, member_id, join_date, verification_status) "+
		"values ($1, $2, $3, $4, $5)",
		username, password, memberId, joinDate, verification)
	if err != nil {
		return err
	}
	return nil
}

func NewMemberAccessRepo(db *sqlx.DB) MemberAccessRepo {
	return &memberAccessRepoImpl{
		db: db,
	}
}
