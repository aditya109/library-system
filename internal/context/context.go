package context

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	logger "github.com/sirupsen/logrus"
)

type Library struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
}

// OpenConnection attempts the connection to the sql db via my
func (l Library) OpenConnection() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s", l.DbUser, l.DbPassword, l.DbHost, l.DbName))
	if err != nil {
		logger.Fatalf("error in opening the connection to the database %s\n", err.Error())
	}
	return db
}

func (l Library) CloseConnection(db *sql.DB) {
	if err := db.Close(); err != nil {
		logger.Fatalf("error in closing connection %s\n", err.Error())
	}
}
