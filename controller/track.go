package controller

import (
	"net/http"
	"satplan/common"
	"satplan/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTrackBySatId(c *gin.Context) {
	satId := c.Query("satid")
	start, err := strconv.ParseInt(c.Query("start"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetRespResult(int(common.FAILED),
			"bad format of start time", nil, 0))
		return
	}
	stop, err := strconv.ParseInt(c.Query("stop"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetRespResult(int(common.FAILED),
			"bad format of stop time", nil, 0))
		return
	}
	track := service.GetSatTrack(satId, start, stop)
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"query success", track, len(*track)))
}
