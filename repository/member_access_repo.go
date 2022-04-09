package repository

import (
	"dating_app_last/model"
	"dating_app_last/util"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type MemberAccessRepo interface {
	Insert(username, password, memberId string, joinDate *time.Time, verification string) error
	Update(id string) error
	Login(username, password string) (*model.MemberAccess, error)
	GetMember(page int) (*[]model.GetMemberResp, error)
}

type memberAccessRepoImpl struct {
	db *sqlx.DB
}

func (m *memberAccessRepoImpl) Login(username, password string) (*model.MemberAccess, error) {
	var memberId, userName, userPassword, verification string
	var joinDate *time.Time
	err := m.db.QueryRow("SElECT * FROM member_access WHERE user_name = $1 AND user_password = $2 AND verification_status = $3",
		username, password, "Y").Scan(&memberId, &userName, &userPassword, &joinDate, &verification)
	if err != nil {
		util.Log.Error().Msg("Unauthorized member")
		return nil, errors.New("Unauthorized member")
	}
	loginReq := new(model.MemberAccess)
	loginReq.UserName = userName
	loginReq.Verification = verification
	loginReq.JoinDate = joinDate
	loginReq.Password = userPassword
	loginReq.MemberId = memberId

	return loginReq, nil
}

func (m *memberAccessRepoImpl) GetMember(page int) (*[]model.GetMemberResp, error) {
	var listMember []model.GetMemberResp
	err := m.db.Select(&listMember, "SELECT user_name, join_date FROM member_access ORDER BY user_name ASC LIMIT 5 OFFSET $1", (5 * (page - 1)))
	if err != nil {
		util.Log.Error().Msg(err.Error())
		return nil, errors.New("Get member error")
	}
	return &listMember, nil
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
