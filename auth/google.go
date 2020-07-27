package auth

import (
	"encoding/json"
	"io/ioutil"

	"github.com/AMYMEME/board-golang/model"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleAuth(authCode string, redirectURI string) (string, error) {
	conf := &oauth2.Config{
		ClientID:     "98250998013-tbbgns4o192oci849p3e8pf7ntfdgl8g.apps.googleusercontent.com",
		ClientSecret: "Z1n3XbXaazNfy_8bTSEcBEwh",
		RedirectURL:  redirectURI,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/plus.login",
		},
		Endpoint: google.Endpoint,
	}

	tok, err := conf.Exchange(oauth2.NoContext, authCode)

	if err != nil {
		err := errors.Wrap(err, "Fail exchange authorization code")
		return "", err
	}

	client := conf.Client(oauth2.NoContext, tok)

	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		err := errors.Wrap(err, "Fail getting google user information")
		return "", err
	}
	defer userInfoResp.Body.Close()

	userInfo, err := ioutil.ReadAll(userInfoResp.Body)

	if err != nil {
		err := errors.Wrap(err, "Fail reading user information response body")
		return "", err
	}

	var googleUserInfo model.GoogleUserInfo

	if err := json.Unmarshal(userInfo, &googleUserInfo); err != nil {
		err := errors.Wrap(err, "Fail user information into binding")
		return "", err
	}

	return googleUserInfo.Email, nil
}
