package service

import (
	"satplan/dao/db"
	"satplan/entity"
)

func GetSenPath(satId string, senName string, start int64, stop int64) *[]entity.SenPath {
	pathInfo := db.FindSenPathInfo(satId, senName)
	minTimeInDb := pathInfo.StartTime
	lastPathPoint := db.FindLastSenPathPoint(satId, senName)
	maxTimeInDb := pathInfo.StartTime + lastPathPoint.TimeOffset

	if maxTimeInDb < stop {
		stop = maxTimeInDb
	}
	if minTimeInDb > start {
		start = minTimeInDb
	}

	start = start - minTimeInDb
	stop = stop - minTimeInDb
	return db.FindSenPath(satId, senName, start, stop)
}

func GetPathPlan(planPara *entity.PlanPara) *[]entity.PathUnit {
	pathUnits := []entity.PathUnit{}

	for _, senId := range *planPara.CheckedSenIds {
		sen, err := db.FindSensorById(senId)
		if err != nil {
			continue
		}
		units := db.FindPathUnit((*sen).SatNoardId,
			(*sen).Name, planPara.Start, planPara.Stop, planPara.Xmin, planPara.Xmax,
			planPara.Ymin, planPara.Ymax)
		pathUnits = append(pathUnits, *units...)
	}

	return &pathUnits
}
