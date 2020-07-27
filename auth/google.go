package auth

import (
	"encoding/json"
	"io/ioutil"

	"github.com/AMYMEME/board-golang/model"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var conf *oauth2.Config

func init() {
	conf = &oauth2.Config{
		ClientID:     "98250998013-tbbgns4o192oci849p3e8pf7ntfdgl8g.apps.googleusercontent.com",
		ClientSecret: "Z1n3XbXaazNfy_8bTSEcBEwh",
		RedirectURL:  "postmessage",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/plus.login",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleAuth(authCode string) (model.UserInfo, error) {
	tok, err := conf.Exchange(oauth2.NoContext, authCode)

	if err != nil {
		err := errors.Wrap(err, "Fail exchange authorization code")
		return model.UserInfo{}, err
	}

	userInfoResp, err := conf.Client(oauth2.NoContext, tok).Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		err := errors.Wrap(err, "Fail getting google user information")
		return model.UserInfo{}, err
	}
	defer userInfoResp.Body.Close()

	userInfo, err := ioutil.ReadAll(userInfoResp.Body)

	if err != nil {
		err := errors.Wrap(err, "Fail reading user information response body")
		return model.UserInfo{}, err
	}

	var googleUserInfo model.UserInfo

	if err := json.Unmarshal(userInfo, &googleUserInfo); err != nil {
		err := errors.Wrap(err, "Fail user information into binding")
		return model.UserInfo{}, err
	}

	return googleUserInfo, nil
}
