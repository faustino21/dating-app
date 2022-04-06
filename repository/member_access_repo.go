package repository

import (
	"dating_app_last/util"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type MemberAccessRepo interface {
	Insert(username, password, memberId string, joinDate *time.Time, verification string) error
	Update(id string) error
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

func (m *memberAccessRepoImpl) Update(id string) error {
	var email *string
	tx := m.db.MustBegin()
	rowsAffected := tx.MustExec("UPDATE member_access SET verification_status = $1 WHERE member_id = $2", "Y", id)
	rows, err := rowsAffected.RowsAffected()
	if rows == 0 {
		util.Log.Error().Msg("Wrong Member Id")
		return err
	}
	err = tx.Get(&email, "SELECT user_name FROM member_access WHERE member_id = $1", id)
	if err != nil {
		util.Log.Error().Msg("Member not registered")
		return errors.New("Member not registered")
	}
	tx.MustExec("INSERT INTO member_contact_information (member_id, email,contact_information_id) VALUES ($1, $2, $3)", id, email, util.GetUuid())
	err = tx.Commit()
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
