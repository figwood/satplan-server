package service

import (
	"satplan/dao/db"
	"satplan/entity"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func GetCurrentUserId(c *gin.Context) int {
	return getSysUser(c).Id
}

func getSysUser(c *gin.Context) *entity.SysUser {
	//email
	email := GetCurrentUserEmail(c)

	//find user info by email
	sysUser := db.FindSysUserByEmail(email)
	if sysUser.Id == 0 {
		log.Debug("cannot find user: " + email)
	}
	return sysUser
}

func CurrentUserIsPlatformAdmin(c *gin.Context) bool {
	sysUser := getSysUser(c)

	return IsPlatformAdmin(sysUser.Id)
}
