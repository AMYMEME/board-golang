package model

type Google struct {
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type GoogleUserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
