package db

import (
	"database/sql"

	// See https://github.com/ryoccd/gochat/log
	logger "github.com/ryoccd/gochat/log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=gochat sslmode=disable")
	if err != nil {
		logger.Error(err)
	}
	return
}
