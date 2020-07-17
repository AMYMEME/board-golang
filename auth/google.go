package auth

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleAuth(authCode string, redirectURI string) (*oauth2.Token, error) {
	fmt.Printf("redirect_uri: %s\n", redirectURI)
	conf := &oauth2.Config{
		ClientID:     "98250998013-tbbgns4o192oci849p3e8pf7ntfdgl8g.apps.googleusercontent.com",
		ClientSecret: "Z1n3XbXaazNfy_8bTSEcBEwh",
		RedirectURL:  redirectURI,
		Scopes: []string{
			"profile",
			"email",
			"https://www.googleapis.com/auth/plus.login",
		},
		Endpoint: google.Endpoint,
	}

	tok, err := conf.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		err := errors.Wrap(err, "Fail converting authorization code into a token.")
		return nil, err
	}

	return tok, nil
}
