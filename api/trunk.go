package api

import (
	"asterisk-monitor/app/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Trunkstatus(ctx *gin.Context) {
	var logs []CallLogModel
	err := mysql.DBOrmInstance.Where("trunk != ''").Find(&logs)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	result := make(map[string]int)
	for _, log := range logs {
		if _, ok := result[log.Trunk]; !ok {
			result[log.Trunk] = 0
		}
		result[log.Trunk] = result[log.Trunk] + 1
	}

	ctx.JSON(http.StatusOK, result)
}
