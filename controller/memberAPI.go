package controller

import (
	"net/http"
	"strconv"

	"github.com/AMYMEME/board-golang/config"
	"github.com/AMYMEME/board-golang/model"
	"github.com/AMYMEME/board-golang/repository"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var logger = config.GetLogger()

func AddMember(c *gin.Context) {
	var request model.Member
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Errorf(errors.Wrap(err, "Fail request body bind to member").Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail request body bind to member"})
		return
	}

	// TODO : 아직 OAuth를 연결 못해서 그냥 AMYEMEME로만 멤버 추가..
	/*
		oauth2
	*/

	autoID := 1
	res := repository.AddMember(autoID, request.Name)
	if res != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res})
		return
	}
	c.JSON(http.StatusOK, gin.H{"member_id": autoID})
}

func GetMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(errors.Wrap(err, "path parameter is invalid"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is invalid"})
		return
	}
	response, err := repository.GetMember(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}
