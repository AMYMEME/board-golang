package controller

import (
	"net/http"
	"strconv"

	"github.com/AMYMEME/board-golang/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AddComment(c *gin.Context) {
	var request model.Comment

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("path parameter is invalid").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorw("ERROR", "body", errors.New("Fail request body bind to comment").Error(),
			"status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to comment"})
		return
	}
	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.PostID = postID
	if err := DB.AddComment(request); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "status_code", http.StatusCreated)
	c.JSON(http.StatusCreated, gin.H{})
}

func DeleteComment(c *gin.Context) {
	logger.Infow("REQUEST", "method", http.MethodDelete, "url", c.Request.URL)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Errorw("ERROR", "body", errors.New("path parameter is invalid").Error(),
			"status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := DB.DeleteComment(id); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("RESPONSE", "status_code", http.StatusOK)
	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("path parameter is invalid").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	var request model.Comment
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("Fail request body bind to comment").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to comment"})
		return
	}
	logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL, "body", request)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := DB.UpdateComment(id, request); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "status_code", http.StatusResetContent)
	c.JSON(http.StatusResetContent, gin.H{})
}
