package api

import (
	"net/http"

	"github.com/AMYMEME/board-golang/auth"
	"github.com/AMYMEME/board-golang/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func userInfoLogic(userInfo model.UserInfo) (string, error) {
	if err := connectionCheck(DB.MyDB); err != nil {
		return "", err
	}

	result, err := DB.CheckProviderInfoExists(userInfo.Provider, userInfo.Email)
	if err != nil {
		return "", err
	}

	if result {
		// exist => getUser
		ID, err := DB.GetProviderInfo(userInfo.Provider, userInfo.Email)
		if err != nil {
			return "", err
		}
		member, err := DB.GetMember(ID)
		if err != nil {
			return "", err
		}

		token, err := auth.CreateToken(member)
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

	token, err := auth.CreateToken(member)
	if err != nil {
		return "", err
	}

	return token, nil
}

func AuthGoogle(c *gin.Context) {
	var request model.GoogleAuth
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorw("ERROR", "body", errors.New("Fail request body bind").Error(), "status_code", http.StatusBadRequest)
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
	var request model.NaverKakaoAuth
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorw("ERROR", "body", errors.New("Fail request body bind").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind"})
		return
	}

	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	if request.Name == "" {
		err := errors.New("user name is necessary")
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	naverUserInfo := model.UserInfo{Provider: "naver", Name: request.Name, Email: request.UniqID}
	token, err := userInfoLogic(naverUserInfo)

	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("RESPONSE", "body", token, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, token)
}

func AuthKakao(c *gin.Context) {
	var request model.NaverKakaoAuth
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorw("ERROR", "body", errors.New("Fail request body bind").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind"})
		return
	}

	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	if request.Name == "" {
		err := errors.New("user name is necessary")
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kakaoUserInfo := model.UserInfo{Provider: "kakao", Name: request.Name, Email: request.UniqID}
	token, err := userInfoLogic(kakaoUserInfo)

	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("RESPONSE", "body", token, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, token)
}
