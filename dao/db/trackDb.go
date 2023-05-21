package db

import (
	"satplan/entity"
)

//=================info===============
func FindSatTrackInfo(satId string) *entity.TrackInfo {
	trackDb := GetSatTrackDb(satId)
	if trackDb == nil {
		return nil
	}

	track := entity.TrackInfo{}
	trackDb.First(&track)
	return &track
}

//=================track===============
func FindLastTrackPoint(satId string) *entity.Track {
	trackDb := GetSatTrackDb(satId)
	if trackDb == nil {
		return nil
	}

	track := entity.Track{}
	trackDb.Order("time_offset desc").First(&track)
	return &track
}

func FindSatTrack(satId string, start int64, stop int64) *[]entity.Track {
	trackDb := GetSatTrackDb(satId)
	if trackDb == nil {
		return nil
	}

	track := []entity.Track{}
	trackDb.Where("time_offset >=? and time_offset <= ?", start, stop).Find(&track)
	return &track
}
