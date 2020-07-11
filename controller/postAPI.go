package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AMYMEME/board-golang/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AddPost(c *gin.Context) {
	var request model.Post
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorw("ERROR", "body", errors.New("Fail request body bind to post").Error(),
			"status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to post"})
		return
	}
	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ID, err := DB.AddPost(request)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", fmt.Sprintf("/post/%d", ID), "status_code", http.StatusOK)
	c.JSON(http.StatusOK, fmt.Sprintf("/post/%d", ID))
}

func GetPost(c *gin.Context) {
	logger.Infow("REQUEST", "method", http.MethodGet, "url", c.Request.URL)

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
	response, err := DB.GetPost(id)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", response, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, response)
}

func GetAllPosts(c *gin.Context) {
	logger.Infow("REQUEST", "method", http.MethodGet, "url", c.Request.URL)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res, err := DB.GetAllPosts()
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", res, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, res)
}

func DeletePost(c *gin.Context) {
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

	if err := DB.DeletePost(id); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("RESPONSE", "status_code", http.StatusOK)
	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("path parameter is invalid").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	var request model.Post
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("Fail request body bind to post").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to post"})
		return
	}
	logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL, "body", request)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := DB.UpdatePost(id, request); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", fmt.Sprintf("/post/%d", id), "status_code", http.StatusOK)
	c.JSON(http.StatusOK, fmt.Sprintf("/post/%d", id))
}
