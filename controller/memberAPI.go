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
		logger.Errorf(errors.Wrap(err, "Fail request body bind to member").Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to member"})
		return
	}
	if err := connectionCheck(DB.MyDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// TODO : 아직 OAuth를 연결 못해서 그냥 AMYEMEME로만 멤버 추가..
	/*
		oauth2
	*/

	autoID := 2
	res := DB.AddMember(autoID, request.Name)
	if res != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("/member/%d", autoID)})
}

func GetMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(errors.Wrap(err, "path parameter is invalid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	if err := connectionCheck(DB.MyDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	response, err := DB.GetMember(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetAllMembers(c *gin.Context) {
	if err := connectionCheck(DB.MyDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	res, err := DB.GetAllMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": res})
}

func DeleteMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(errors.Wrap(err, "path parameter is invalid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	if err := connectionCheck(DB.MyDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := DB.DeleteMember(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(errors.Wrap(err, "path parameter is invalid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}

	var request model.Member
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Errorf(errors.Wrap(err, "Fail request body bind to member").Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to member"})
		return
	}

	if err := connectionCheck(DB.MyDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := DB.UpdateMember(id, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("/member/%d", id)})
}
