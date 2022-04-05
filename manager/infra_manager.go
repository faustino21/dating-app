package manager

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type InfraManager interface {
	GetSqlConn() *sqlx.DB
}

type infraManager struct {
	infra *sqlx.DB
}

func (i *infraManager) GetSqlConn() *sqlx.DB {
	return i.infra
}

func NewInfra(dataSourceName string) InfraManager {
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &infraManager{
		infra: conn,
	}
}
