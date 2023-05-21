package db

import (
	"fmt"
	"satplan/entity"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetSenPathDb(satId string, senName string) *gorm.DB {
	dbName := fmt.Sprintf("%s/%s/%s.sqlite", dataFolder, satId, senName)
	pathDb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Errorf("error opening db file: %s", dbName)
		log.Error(err)
		return nil
	}

	return pathDb
}

func GetSatTrackDb(satId string) *gorm.DB {
	dbName := fmt.Sprintf("%s/%s.sqlite", dataFolder, satId)
	trackDb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Errorf("error opening db file: %s", dbName)
		log.Error(err)
		return nil
	}

	return trackDb
}

//=================sysuser===============

func FindAllSysUsers() *[]entity.SysUser {
	userList := []entity.SysUser{}
	query := satDb.Find(&userList)
	if query.Error != nil {
		log.Error(query.Error)
	}

	return &userList
}

func FindSysUserByEmail(email string) *entity.SysUser {
	user := entity.SysUser{}
	satDb.Where("email= ? ", email).First(&user)
	return &user
}

func FindSysUserByUserId(userId int) *entity.SysUser {
	user := entity.SysUser{}
	satDb.Where("id= ? ", userId).First(&user)
	return &user
}

func FindPrivilegeList(userId int) *[]entity.PrivilegeMenuVO {
	/*
		privilegeList := []entity.PrivilegeMenuVO{}
		sql := "select distinct m.menu_id as id,m.p_id,m.menu_name,m.url,m.icon,m.com_name,m.sort," +
			"m.routing_type from sys_menu m, sys_privilege p, sys_user u where u.role_id=p.role_id " +
			"and p.menu_id=m.menu_id and u.id=? order by m.sort"
		query := satDb.Raw(sql, userId).Scan(&privilegeList)
		if query.Error != nil {
			log.Debug("error: " + query.Error.Error())
		}
		return &privilegeList
	*/
	return &[]entity.PrivilegeMenuVO{
		{
			Id:  1,
			PId: 1,
			Url: "/user",
		},
		{
			Id:  2,
			PId: 1,
			Url: "/satellite",
		},
	}
}

//=================sensor===============
func DeleteSensors() {
	satDb.Delete(entity.Sensor{}, "1=1")
}

func DeleteSensorById(id int) error {
	query := satDb.Delete(entity.Sensor{}, "id= ?", id)
	return query.Error
}

func DeleteSensorBySatIdAndName(satId string, senName string) error {
	query := satDb.Delete(entity.Sensor{},
		"sat_noard_id= ? and name =?", satId, senName)
	return query.Error
}

func CreateSensor(sen *entity.Sensor) error {
	query := satDb.Create(sen)
	return query.Error
}

func BatCreateSensor(sens *[]entity.Sensor) error {
	query := satDb.CreateInBatches(sens, len(*sens))
	return query.Error
}

func FindAllSensors() *[]entity.Sensor {
	senList := []entity.Sensor{}
	query := satDb.Find(&senList)
	if query.Error != nil {
		log.Error(query.Error)
	}

	return &senList
}

func FindSensorById(id int) (*entity.Sensor, error) {
	sen := entity.Sensor{}
	query := satDb.Where("id= ? ", id).First(&sen)
	return &sen, query.Error
}

func FindSensorBySatId(noardId string) (*[]entity.Sensor, error) {
	senList := []entity.Sensor{}
	query := satDb.Where("sat_noard_id= ? ", noardId).Find(&senList)
	return &senList, query.Error
}

func FindSensorBySatIdAndName(noardId string, senName string) (*entity.Sensor, error) {
	sen := entity.Sensor{}
	query := satDb.Where("sat_noard_id= ? and name= ?", noardId, senName).First(&sen)
	return &sen, query.Error
}

func SaveSensor(sen *entity.Sensor) error {
	query := satDb.Save(sen)
	return query.Error
}

//=================satellite===============

func FindAllSatellites() *[]entity.Satellite {
	satList := []entity.Satellite{}
	query := satDb.Find(&satList)
	if query.Error != nil {
		log.Error(query.Error)
	}

	return &satList
}

func DeleteSatelliteById(id int) error {
	query := satDb.Delete(entity.Satellite{}, "id= ?", id)
	return query.Error
}

func DeleteSatellites() {
	satDb.Delete(entity.Satellite{}, "1=1")
}

func FindSatelliteById(id int) (*entity.Satellite, error) {
	sat := entity.Satellite{}
	query := satDb.Where("id= ? ", id).First(&sat)
	return &sat, query.Error
}

func FindSatelliteByNoardId(noardId string) (*entity.Satellite, error) {
	sat := entity.Satellite{}
	query := satDb.Where("noard_id= ? ", noardId).First(&sat)
	return &sat, query.Error
}

func CreateSatellite(sat *entity.Satellite) error {
	query := satDb.Create(sat)
	return query.Error
}

func BatCreateSatellite(sats *[]entity.Satellite) error {
	query := satDb.CreateInBatches(sats, len(*sats))
	return query.Error
}

func SaveSatellite(sat *entity.Satellite) error {
	query := satDb.Save(sat)
	return query.Error
}

//=================tle===============
func DeleteTles() {
	satDb.Delete(entity.Tle{}, "1=1")
}

func CreateTle(tle *entity.Tle) error {
	query := satDb.Create(tle)
	return query.Error
}

func BatCreateTle(tles *[]entity.Tle) error {
	query := satDb.CreateInBatches(tles, len(*tles))
	return query.Error
}
