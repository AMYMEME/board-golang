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

type NaverAuth struct {
	AccessToken string `json:"access_token"`
	State       string `json:"state"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}

type NaverUserInfo struct {
	Resultcode string                `json:"resultcode"`
	Message    string                `json:"message"`
	Response   NaverUserInfoResponse `json:"response"`
}

type NaverUserInfoResponse struct {
	Email        string `json:"email"`
	Nickname     string `json:"nickname"`
	ProfileImage string `json:"profile_image"`
	Age          string `json:"age"`
	Gender       string `json:"gender"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Birthday     string `json:"birthday"`
}
