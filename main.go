package main

import (
	"fmt"
	"satplan/common"
	"satplan/dao/db"
	"satplan/syscfg"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		//DisableColors: true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	log.SetReportCaller(true)

	//log level
	strLogLevel := common.GetEnvValue("LOG_LEVEL", "DEBUG")
	logLevel := log.DebugLevel
	switch strings.ToLower(strLogLevel) {
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "warn":
		logLevel = log.WarnLevel
	case "error":
		logLevel = log.ErrorLevel
	case "fata":
		logLevel = log.FatalLevel
	}
	log.SetLevel(logLevel)

	defer db.Close()
	httpPort := 8080

	//auth and router
	authMiddleware := syscfg.NewAuthMiddleware()
	router := syscfg.NewGinRouterWithAuth(authMiddleware)

	log.Infof("start listening on %d", httpPort)
	router.Run(fmt.Sprintf(":%d", httpPort))
}
