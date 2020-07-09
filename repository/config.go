package repository

import (
	"database/sql"
	"github.com/AMYMEME/board-golang/config"
)

var DB, _ = sql.Open("mysql", "root:root@(127.0.0.1:3306)/board")
var Logger = config.GetLogger()

