package pkg

import (
	"fmt"
	"main/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type pgDatabase struct{}

func NewPgDatabase() *pgDatabase {
	return &pgDatabase{}
}

func (db *pgDatabase) ConnectPgDatabase(p utils.PostgresConf) *sqlx.DB {
	if !p.Enable {
		return nil
	}

	strConn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Database, p.Password, p.SslMode)
	var err error
	dbInstance, err := sqlx.Open("postgres", strConn)
	dbInstance.SetMaxIdleConns(int(p.MaxIdleConn))
	dbInstance.SetMaxOpenConns(int(p.MaxOpenConn))
	if err != nil {
		panic(err)
	}

	err = dbInstance.Ping()
	if err != nil {
		panic(err)
	}

	utils.ShowInfoLogs("Connect to postgres database %s:%d/%s successful", p.Host, p.Port, p.Database)
	return dbInstance
}
