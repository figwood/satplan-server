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

func AddSensor(c *gin.Context) {
	var sensorInDTO entity.NewSensorInDTO
	c.ShouldBindBodyWith(&sensorInDTO, binding.JSON)
	currentUserId := service.GetCurrentUserId(c)
	if !service.IsPlatformAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"method not allowed", nil, 0))
		return
	}

	sensorId, err := service.AddSensor(&sensorInDTO)
	if err != nil {
		log.Debug(err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", []int{sensorId}, 1))
	}
}

func GetAllSensors(c *gin.Context) {
	sensors := service.GetAllSensors()
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"success", sensors, len(*sensors)))
}

func GetSensorBySatId(c *gin.Context) {
	satId := c.Query("satid")
	sensors, err := service.GetSensorBySatId(satId)
	if err != nil {
		log.Debug("GetSensorGroups: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", *sensors, len(*sensors)))
	}
}

func GetSensorById(c *gin.Context) {
	sensorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.FAILED),
			"bad param for sensor id", nil, 0))
	}
	sensor := service.GetSensorById(sensorId)
	if sensor.Id == 0 {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"failed", nil, 0))
	} else {
		sensors := []entity.Sensor{*sensor}
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", &sensors, 1))
	}
}

func DeleteSensor(c *gin.Context) {
	sensorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.FAILED),
			"bad param for sensor id", nil, 0))
	}
	err = service.DeleteSensorById(sensorId)
	if err != nil {
		log.Debug("DeleteSensor: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", []bool{true}, 1))
	}
}

func UpdateSensor(c *gin.Context) {
	sensorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.FAILED),
			"bad param for sensor id", nil, 0))
	}
	var sensorInDTO entity.SensorDTO
	c.ShouldBindBodyWith(&sensorInDTO, binding.JSON)

	err = service.UpdateSensor(sensorId, &sensorInDTO)
	if err != nil {
		log.Debug("UpdateSensor: " + err.Error())
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			err.Error(), nil, 0))
	} else {
		c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
			"success", []bool{true}, 1))
	}
}
