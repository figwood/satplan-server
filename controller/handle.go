package controller

import (
	"fmt"
	"net/http"
	"satplan/common"
	"time"

	"github.com/gin-gonic/gin"
)

func TempTest(c *gin.Context) {
	data := "ok"
	//service.UpdateTles()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func HelloGin(c *gin.Context) {
	t := time.Now()
	res := fmt.Sprintf("%s: welcome, current version: %v",
		t.Format("2006-01-02 15:04:05"), common.APP_VERSION)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, common.GetRespResult(int(common.SUCCEED),
		"got version", []string{common.APP_VERSION}, 1))
}
