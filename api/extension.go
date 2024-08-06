package api

import (
	"asterisk-monitor/app/ami"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Extensionstatus(ctx *gin.Context) {
	data, err := ami.ExtnesionStatus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// UNKNOWN
	// NOT_INUSE
	// INUSE
	// BUSY
	// INVALID
	// UNAVAILABLE
	// RINGING
	// RINGINUSE
	// ONHOLD
	result := map[string]int{"Busy": 0, "Unavailable": 0, "Not in use": 0, "Ringing": 0, "In use": 0, "On Hold": 0}
	for _, event := range data {
		// 忽略软电话
		if strings.Contains(event.ObjectName, "app_") {
			continue
		}
		switch event.DeviceState {
		case "Busy":
			result["Busy"] += 1
		case "Unavailable":
			result["Unavailable"] += 1
		case "Not in use":
			result["Not in use"] += 1
		case "Ringing":
			result["Ringing"] += 1
		case "In use":
			result["In use"] += 1
		case "On Hold":
			result["On Hold"] += 1
		}
	}
	ctx.JSON(http.StatusOK, result)
}
