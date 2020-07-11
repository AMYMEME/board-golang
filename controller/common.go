package controller

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/AMYMEME/board-golang/config"

	"github.com/AMYMEME/board-golang/repository"
)

var logger = config.GetLogger()

var DB = repository.DBConfig{
	MyDB: setupDB(),
}

func setupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/board?parseTime=true")
	if err != nil {
		logger.Errorf("Fail DB connection")
	}

	return db
}

func connectionCheck(db *sql.DB) error {
	if db == nil {
		err := errors.New("Fail sql driver connection")
		return err
	}
	return nil
}
