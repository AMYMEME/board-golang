package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AMYMEME/board-golang/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AddMember(c *gin.Context) {
	var request model.Member
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL)
		logger.Errorw("ERROR", "body", errors.New("Fail request body bind to member").Error(),
			"status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to member"})
		return
	}
	logger.Infow("REQUEST", "method", http.MethodPost, "url", c.Request.URL, "body", request)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO : 아직 OAuth를 연결 못해서 그냥 AMYEMEME로만 멤버 추가..

	autoID := 1
	ID, err := DB.AddMember(autoID, request.Name)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", fmt.Sprintf("/member/%d", ID), "status_code", http.StatusOK)
	c.JSON(http.StatusOK, fmt.Sprintf("/member/%d", ID))
}

func GetMember(c *gin.Context) {
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
	response, err := DB.GetMember(id)
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", response, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, response)
}

func GetAllMembers(c *gin.Context) {
	logger.Infow("REQUEST", "method", http.MethodGet, "url", c.Request.URL)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res, err := DB.GetAllMembers()
	if err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", res, "status_code", http.StatusOK)
	c.JSON(http.StatusOK, res)
}

func DeleteMember(c *gin.Context) {
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

	if err := DB.DeleteMember(id); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("RESPONSE", "status_code", http.StatusOK)
	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("path parameter is invalid").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	var request model.Member
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL)
		logger.Errorf("ERROR", "body", errors.New("Fail request body bind to member").Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to member"})
		return
	}
	logger.Infow("REQUEST", "method", http.MethodPut, "url", c.Request.URL, "body", request)

	if err := connectionCheck(DB.MyDB); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := DB.UpdateMember(id, request); err != nil {
		logger.Errorw("ERROR", "body", err.Error(), "status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("RESPONSE", "body", fmt.Sprintf("/member/%d", id), "status_code", http.StatusOK)
	c.JSON(http.StatusOK, fmt.Sprintf("/member/%d", id))
}
