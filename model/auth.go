package model

type GoogleAuth struct {
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type UserInfo struct {
	Provider string `json:"provider"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
