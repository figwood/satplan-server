package db

import (
	"satplan/common"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	log "github.com/sirupsen/logrus"
)

var satDb *gorm.DB
var dataFolder = "../data"

func Close() {
	sqlDB, _ := satDb.DB()
	sqlDB.Close()
}

func init() {
	var err error
	dataFolder = common.GetEnvValue("DATA_FOLDER", dataFolder)
	dbName := dataFolder + "/sat.db"

	// github.com/mattn/go-sqlite3
	satDb, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	log.Info("database inited successfully")
}
