package service

import (
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func init() {
	InitCron()
}

func InitCron() {
	c = cron.New()
	//c.AddFunc("@every 11m", UpdateJenkinsOverview)
	//c.AddFunc("@every 1h", SyncGitProjects)
	//every two days at 11:00 pm
	//c.AddFunc("0 23 */2 * *", UpdateToolVersions)
	c.AddFunc("0 0 * * *", CronUpdateTLE)
	c.Start()
}

func CronUpdateTLE() {
	UpdateTles()
}
