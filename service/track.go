package service

import (
	"satplan/dao/db"
	"satplan/entity"
)

func GetSatTrack(satId string, start int64, stop int64) *[]entity.Track {
	trackInfo := db.FindSatTrackInfo(satId)
	minTimeInDb := trackInfo.StartTime
	lastTrackPoint := db.FindLastTrackPoint(satId)
	maxTimeInDb := trackInfo.StartTime + lastTrackPoint.TimeOffset

	if maxTimeInDb < stop {
		stop = maxTimeInDb
	}
	if minTimeInDb > start {
		start = minTimeInDb
	}

	start = start - minTimeInDb
	stop = stop - minTimeInDb
	return db.FindSatTrack(satId, start, stop)
}
