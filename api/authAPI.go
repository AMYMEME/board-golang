package api

import (
	"net/http"

	"github.com/AMYMEME/board-golang/model"

	"github.com/AMYMEME/board-golang/auth"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AuthGoogle(c *gin.Context) {
	var request model.Google
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("Fail request body bind").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind"})
		return
	}

	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	token, err := auth.GoogleAuth(request.Code, request.RedirectURI)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
