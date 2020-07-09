package repository

import (
	"github.com/AMYMEME/board-golang/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func AddMember(autoID int, name string) error {
	if DB == nil {
		err := errors.New("Fail sql driver connection")
		Logger.Errorf(err.Error())
		return err
	}

	_, err := DB.Exec("INSERT INTO board.member (id, name, level) VALUES (?, ?, ?)", autoID, name, 0)
	if err != nil {
		err := errors.Wrap(err, "Fail sql query")
		Logger.Errorf(err.Error())
		return err
	}
	return nil
}

func GetMember(ID int) (model.Member, error) {
	if DB == nil {
		err := errors.New("Fail sql driver connection")
		Logger.Errorf(err.Error())
		return model.Member{}, err
	}

	var member model.Member
	err := DB.QueryRow("SELECT id, name, level FROM board.member WHERE id = ?", ID).Scan(&member.ID, &member.Name, &member.Level)

	if err != nil {
		err := errors.Wrap(err, "There is no such member")
		Logger.Errorf(err.Error())
		return model.Member{}, err
	}
	return member, nil
}
