package repository

import (
	"fmt"

	"github.com/AMYMEME/board-golang/model"
	"github.com/pkg/errors"
)

func (d *DBConfig) AddProviderInfo(oauth model.Oauth) (int, error) {
	res, err := d.MyDB.Exec("INSERT INTO board.oauth (provider, provider_id) VALUES (?, ?)", oauth.Provider, oauth.ProviderID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return 0, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		err := errors.Wrap(err, "Fail getting added row's id")
		return 0, err
	}
	return int(ID), nil
}

func (d *DBConfig) GetProviderInfo(provider string, providerID string) (int, error) {
	var ID int
	err := d.MyDB.QueryRow("SELECT id FROM board.oauth WHERE provider = ? AND provider_id = ?", provider, providerID).
		Scan(&ID)

	if err != nil {
		err := errors.Wrap(err, "There is no such provider info")
		return 0, err
	}
	return ID, nil
}

func (d *DBConfig) CheckProviderInfoExists(provider string, providerID string) (bool, error) {
	var result int
	err := d.MyDB.QueryRow("SELECT EXISTS (SELECT * FROM board.oauth WHERE provider = ? AND provider_id = ?) AS SUCCESS", provider, providerID).
		Scan(&result)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return false, err
	}
	fmt.Println("result : ", result)
	switch result {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return false, err
	}
}
