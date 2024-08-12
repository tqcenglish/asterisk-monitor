package api

import (
	"asterisk-monitor/app/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SystemInfoResp struct {
	SystemUptime int    `json:"systemUptime"`
	VoipUptime   int    `json:"voipUptime"`
	VoipReload   int    `json:"voipReload"`
	WebUptime    int    `json:"webUptime"`
	Version      string `json:"version"`
}

type StorageInfoResp struct {
	Record    int64 `json:"record"`
	VoiceMail int64 `json:"voicemail"`
	System    int64 `json:"system"`
	Music     int64 `json:"music"`
	Other     int64 `json:"other"`
}

func SystemInfo(c *gin.Context) {
	asteriksUptimeSeconds, err := utils.RunCmd(`asterisk -rx 'core show uptime seconds'`)
	if err != nil {
		logrus.Error(err)
	}
	var asteriskUptime int
	var asteriskReload int
	for _, timeStr := range strings.Split(string(asteriksUptimeSeconds), "\n") {
		if timeStr != "" && strings.Contains(timeStr, "System uptime:") {
			asteriskUptime, _ = strconv.Atoi(strings.Split(timeStr, ": ")[1])
		}
		if timeStr != "" && strings.Contains(timeStr, "Last reload:") {
			asteriskReload, _ = strconv.Atoi(strings.Split(timeStr, ": ")[1])
		}
	}

	systemUptime, err := utils.RunCmd(`uptime -s`)
	if err != nil {
		logrus.Error(err)
	}

	duration, err := time.Parse("2006-01-02 15:04:05", strings.TrimSuffix(string(systemUptime), "\n"))
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("123", asteriskReload, asteriskUptime, duration.String())
	c.JSON(http.StatusOK, SystemInfoResp{
		SystemUptime: int(time.Since(duration).Hours()),
		VoipUptime:   asteriskUptime / 60 / 60,
		VoipReload:   asteriskReload / 60 / 60,
		Version:      "1.0.0",
	})
}

func StorageInfo(c *gin.Context) {
	// du -sh /var/spool/asterisk/recording
	// du -sh /
	c.JSON(http.StatusOK, StorageInfoResp{
		Record:    93,
		System:    32,
		Music:     44,
		VoiceMail: 65,
		Other:     22,
	})
}
