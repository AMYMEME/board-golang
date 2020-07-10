package repository

import (
	"github.com/AMYMEME/board-golang/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func (d *DBConfig) AddMember(autoID int, name string) error {
	_, err := d.MyDB.Exec("INSERT INTO board.member (id, name, level) VALUES (?, ?, ?)", autoID, name, 0)
	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		logger.Errorf(err.Error())
		return err
	}
	return nil
}

func (d *DBConfig) GetMember(ID int) (model.Member, error) {
	var member model.Member
	err := d.MyDB.QueryRow("SELECT id, name, level FROM board.member WHERE id = ?", ID).Scan(&member.ID, &member.Name, &member.Level)

	if err != nil {
		err := errors.Wrap(err, "There is no such member")
		logger.Errorf(err.Error())
		return model.Member{}, err
	}
	return member, nil
}

func (d *DBConfig) GetAllMembers() ([]model.Member, error) {
	rows, err := d.MyDB.Query("SELECT id, name, level FROM board.member")

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		logger.Errorf(err.Error())
		return nil, err
	}
	var results []model.Member
	defer rows.Close()

	for rows.Next() {
		var member model.Member
		err := rows.Scan(&member.ID, &member.Name, &member.Level)
		if err != nil {
			err := errors.Wrap(err, "Fail loading sql results")
			logger.Errorf(err.Error())
			return nil, err
		}
		results = append(results, member)
	}
	return results, nil
}

func (d *DBConfig) DeleteMember(ID int) error {
	res, err := d.MyDB.Exec("DELETE FROM board.member WHERE id = ?", ID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		logger.Errorf(err.Error())
		return err
	}

	n, _ := res.RowsAffected()
	if n != 1 {
		err := errors.Wrap(err, "There is no such member")
		logger.Errorf(err.Error())
		return err
	}
	return nil
}

func (d *DBConfig) UpdateMember(ID int, newMemberInfo model.Member) error {
	res, err := d.MyDB.Exec("UPDATE board.member SET name = ?, level = ? WHERE id = ?",
		newMemberInfo.Name, newMemberInfo.Level, ID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		logger.Errorf(err.Error())
		return err
	}

	n, _ := res.RowsAffected()
	if n != 1 {
		err := errors.Wrap(err, "There is no such member")
		logger.Errorf(err.Error())
		return err
	}
	return nil
}
