package controller

import (
	"net/http"
	"satplan/common"
	"satplan/service"
	"strconv"

	log "github.com/sirupsen/logrus"

	entity "satplan/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddSatellite(c *gin.Context) {
	var newSat entity.NewSatDTO
	c.ShouldBindBodyWith(&newSat, binding.JSON)

	currentUserId := service.GetCurrentUserId(c)
	if !service.IsPlatformAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"method not allowed", nil, 0))
		return
	}
	satelliteId, err := service.AddSatellite(&newSat)
	if err != nil {
		log.Debug(err.Error())
		c.JSON(http.StatusInternalServerError,
			common.GetRespResult(int(common.FAILED), err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", []int{satelliteId}, 1))
	}
}

func GetAllSatellites(c *gin.Context) {
	satellites := service.GetAllSatellites()
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"success", satellites, len(*satellites)))
}

func GetSatTree(c *gin.Context) {
	satellites := service.GetSatTree()
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"success", satellites, len(*satellites)))
}

func GetSatelliteById(c *gin.Context) {
	satId := c.Param("id")

	satellite, err := service.GetSatelliteById(satId)
	if err != nil {
		log.Debug("GetSatelliteById: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", []entity.Satellite{*satellite}, 1))
	}
}

func UpdateSatellite(c *gin.Context) {
	var satDTO entity.SatDTO
	c.ShouldBindBodyWith(&satDTO, binding.JSON)
	currentUserId := service.GetCurrentUserId(c)
	if !service.IsPlatformAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"method not allowed", nil, 0))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"bad satellite id", nil, 0))
		return
	}
	err = service.UpdateSatellite(id, &satDTO)
	if err != nil {
		log.Debug("UpdateSatellite: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "success", nil, 1))
	}
}

func DeleteSatellite(c *gin.Context) {
	currentUserId := service.GetCurrentUserId(c)
	if !service.IsPlatformAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"method not allowed", nil, 0))
		return
	}
	satId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	}
	err = service.DeleteSatelliteById(satId)
	if err != nil {
		log.Debug("DeleteSatellite: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", []bool{true}, 1))
	}
}

func UpdateTles(c *gin.Context) {
	currentUserId := service.GetCurrentUserId(c)
	if !service.IsPlatformAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"method not allowed", nil, 0))
		return
	}
	err := service.UpdateTles()
	if err != nil {
		log.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "success", nil, 1))
	}
}

func RecalPath(c *gin.Context) {
	currentUserId := service.GetCurrentUserId(c)
	if !service.IsPlatformAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"method not allowed", nil, 0))
		return
	}
	err := service.RecalPath()
	if err != nil {
		log.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED), "success", nil, 1))
	}
}
