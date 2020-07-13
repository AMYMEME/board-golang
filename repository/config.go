package repository

import (
	"database/sql"

	"github.com/AMYMEME/board-golang/config"
)

type DBConfig struct {
	MyDB *sql.DB
}

var logger = config.GetLogger()
