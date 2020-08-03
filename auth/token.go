package auth

import (
	"github.com/AMYMEME/board-golang/model"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(member model.Member) (string, error) {
	mySigningKey := []byte("board-golang")

	type MyCustomClaims struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		member.ID,
		member.Name,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "board-golang",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
