package repository

import (
	"time"

	"github.com/pkg/errors"
)

func AddPost(memberID string, body string) {
	_, err := DB.Exec("INSERT INTO board.post (member_id, body, datetime) VALUES (?, ?, ?)", memberID, body, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		Logger.Errorf(errors.Wrap(err, "Fail sql query").Error())
	}
}
