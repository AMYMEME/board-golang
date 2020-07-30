package auth

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/AMYMEME/board-golang/model"
)

func NaverAuth(tokenType string, accessToken string) (model.UserInfo, error) {
	getInfoURL := `https://openapi.naver.com/v1/nid/me`

	req, err := http.NewRequest(http.MethodGet, getInfoURL, nil)
	if err != nil {
		err := errors.Wrap(err, "Fail making HTTP request for naver API")
		return model.UserInfo{}, err
	}
	req.Header.Add("Authorization", tokenType+" "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err := errors.Wrap(err, "Fail sending HTTP request to naver API")
		return model.UserInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := errors.Wrap(err, "Fail sending HTTP request to naver API")
		return model.UserInfo{}, err
	}

	var naverUserInfo model.NaverUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&naverUserInfo); err != nil {
		err := errors.Wrap(err, "Fail binding HTTP response from naver API")
		return model.UserInfo{}, err
	}

	return model.UserInfo{Provider: "naver", Name: naverUserInfo.Response.Name, Email: naverUserInfo.Response.ID}, nil
}
