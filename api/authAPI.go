package api

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/AMYMEME/board-golang/auth"
	"github.com/AMYMEME/board-golang/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func createToken(member model.Member) (string, error) {
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

func userInfoLogic(provider string, providerID string) (string, error) {
	if err := connectionCheck(DB.MyDB); err != nil {
		return "", err
	}

	if DB.CheckProviderInfoExists(provider, providerID) {
		// exist => getUser
		ID, err := DB.GetProviderInfo(provider, providerID)
		if err != nil {
			return "", err
		}
		member, err := DB.GetMember(ID)
		if err != nil {
			return "", err
		}

		token, err := createToken(member)
		if err != nil {
			return "", err
		}

		return token, nil
	}
	// new register

}

func AuthGoogle(c *gin.Context) {
	var request model.Google
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("Fail request body bind").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind"})
		return
	}

	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	email, err := auth.GoogleAuth(request.Code, request.RedirectURI)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := userInfoLogic("google", email)

	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
