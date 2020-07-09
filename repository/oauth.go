package repository

import (
	"database/sql"

	"github.com/pkg/errors"
)

func AddProviderInfo(provider string, providerID string) int {
	res, err := DB.Exec("INSERT INTO board.oauth (provider, provider_id) VALUES (?, ?)", provider, providerID)
	if err != nil {
		Logger.Errorf(errors.Wrap(err, "Fail sql query").Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		Logger.Errorf(errors.Wrap(err, "Fail getting added row's id").Error())
	}
	defer DB.Close()
	return int(id)
}

func GetProviderInfo(provider string, providerID string) (int, error) {
	if DB == nil {
		err := errors.New("Fail sql driver connection")
		Logger.Errorf(err.Error())
		return 0, err
	}
	var autoID int
	err := DB.QueryRow("SELECT id FROM board.oauth WHERE provider = ? and provider_id = ?", provider, providerID).Scan(&autoID)
	if err == sql.ErrNoRows {
		err := errors.New("There is no such Provider Info")
		Logger.Errorf(err.Error())
		return 0, err
	}
	defer DB.Close()
	return autoID, nil
}
