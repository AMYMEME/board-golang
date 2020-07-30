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

func userInfoLogic(userInfo model.UserInfo) (string, error) {
	if err := connectionCheck(DB.MyDB); err != nil {
		return "", err
	}

	if DB.CheckProviderInfoExists(userInfo.Provider, userInfo.Email) {
		// exist => getUser
		ID, err := DB.GetProviderInfo(userInfo.Provider, userInfo.Email)
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

	ID, err := DB.AddProviderInfo(model.Oauth{
		Provider:   userInfo.Provider,
		ProviderID: userInfo.Email,
	})
	if err != nil {
		return "", err
	}

	if err := DB.AddMember(ID, userInfo.Name); err != nil {
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

func AuthGoogle(c *gin.Context) {
	var request model.GoogleAuth
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("Fail request body bind").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind"})
		return
	}

	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	googleUserInfo, err := auth.GoogleAuth(request.Code)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	googleUserInfo.Provider = "google"
	token, err := userInfoLogic(googleUserInfo)

	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", token, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, token)
}

func AuthNaver(c *gin.Context) {
	var request model.NaverAuth
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("Fail request body bind").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind"})
		return
	}

	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	naverUserInfo, err := auth.NaverAuth(request.TokenType, request.AccessToken)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := userInfoLogic(naverUserInfo)

	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("RESPONSE", "body", token, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, token)
}
