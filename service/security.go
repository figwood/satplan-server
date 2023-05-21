package service

import (
	"errors"
	"fmt"
	"satplan/common"
	"satplan/dao/db"
	"satplan/entity"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetCurrentUserEmail(c *gin.Context) string {
	//get email
	claims := jwt.ExtractClaims(c)
	return fmt.Sprintf("%s", claims["mail"])
}

func GetUserInfo(c *gin.Context) (entity.CurrentUserInfo, error) {
	//get email
	email := GetCurrentUserEmail(c)

	//find user info by email
	sysUser := db.FindSysUserByEmail(email)
	if sysUser.Id == 0 {
		return entity.CurrentUserInfo{}, errors.New("cannot find user: " + email)
	}

	//查找可访问的菜单
	privilegeMenuVOS := db.FindPrivilegeList(sysUser.Id)
	menuVOS := filterMenuVO(privilegeMenuVOS, func(p entity.PrivilegeMenuVO) bool {
		return len(p.Url) != 0
	})
	cui := entity.CurrentUserInfo{
		Id:     sysUser.Id,
		Name:   sysUser.UserName,
		RoleId: 1,
		//RoleList:   orgRoleMap,
		//AdminId:  sysUser.AdminId,
		MenuList: menuVOS,
		//ButtonList: buttonVOS,
	}

	return cui, nil
}

func IsPlatformAdmin(userID int) bool {
	return userID == int(common.PLATFORM_ADMIN)
}

func filterMenuVO(vs *[]entity.PrivilegeMenuVO, f func(entity.PrivilegeMenuVO) bool) []entity.PrivilegeMenuVO {
	vsf := make([]entity.PrivilegeMenuVO, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
