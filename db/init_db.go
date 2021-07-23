package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	// See https://github.com/ryoccd/gochat/log
	logger "github.com/ryoccd/gochat/log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=[USERNAME] password=[PASSWORD] dbname=gochat sslmode=disable")
	if err != nil {
		logger.Error(err)
	}
	return
}
