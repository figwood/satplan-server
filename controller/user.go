package controller

import (
	"net/http"
	"satplan/common"
	"satplan/service"

	entity "satplan/entity"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userInfo, _ := service.GetUserInfo(c)
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"success", []entity.CurrentUserInfo{userInfo}, 1))
}

func GetAllUsers(c *gin.Context) {
	currentUserId := service.GetCurrentUserId(c)
	//权限判断，需要管理员
	if !service.IsPlatformAdmin(currentUserId) {
		c.JSON(http.StatusInternalServerError, common.GetRespResult(int(common.FAILED),
			"method not allowed", nil, 0))
		return
	}

	users := service.GetAllUsers()
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"success", *users, len(*users)))
}
