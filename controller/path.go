package controller

import (
	"net/http"
	"satplan/common"
	entity "satplan/entity"
	"satplan/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetPathBySenId(c *gin.Context) {
	senName := c.Query("senname")
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
	senPath := service.GetSenPath(satId, senName, start, stop)
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"query success", senPath, len(*senPath)))
}

func GetPathPlan(c *gin.Context) {
	var planPara entity.PlanPara
	c.ShouldBindBodyWith(&planPara, binding.JSON)

	senPath := service.GetPathPlan(&planPara)
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"query success", senPath, len(*senPath)))
}
