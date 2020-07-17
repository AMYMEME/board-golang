package model

import "time"

type Member struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type Post struct {
	ID       int       `json:"id"`
	MemberID int       `json:"member_id"`
	Title    string    `json:"title"`
	Contents string    `json:"contents"`
	Datetime time.Time `json:"datetime"`
}

type Oauth struct {
	ID         int    `json:"id"`
	Provider   string `json:"provider"`
	ProviderID string `json:"provider_id"`
}

type Comment struct {
	ID       int       `json:"id"`
	PostID   int       `json:"post_id"`
	MemberID int       `json:"member_id"`
	Contents string    `json:"contents"`
	Datetime time.Time `json:"datetime"`
}

type Google struct {
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}
