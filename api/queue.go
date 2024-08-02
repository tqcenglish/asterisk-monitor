package api

import (
	"asterisk-monitor/app/ami"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueueStatus(ctx *gin.Context) {
	data1, data2, err := ami.QueueStatus()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"params":   data1,
		"memebers": data2,
	})
}
