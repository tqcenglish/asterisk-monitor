package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemInfoResp struct {
	SystemUptime int    `json:"systemUptime"`
	VoipUptime   int    `json:"voipUptime"`
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
	c.JSON(http.StatusOK, SystemInfoResp{
		SystemUptime: 400,
		VoipUptime:   200,
		WebUptime:    130,
		Version:      "1.0.0",
	})
}

func StorageInfo(c *gin.Context) {
	c.JSON(http.StatusOK, StorageInfoResp{
		Record:    93,
		System:    32,
		Music:     44,
		VoiceMail: 65,
		Other:     22,
	})
}
