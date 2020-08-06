package model

type UserInfo struct {
	Provider string `json:"provider"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type ProviderInfo struct {
	Name   string `json:"name"`
	UniqID string `json:"uniq_id"`
}
