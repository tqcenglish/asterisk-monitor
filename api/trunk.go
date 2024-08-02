package api

import (
	"asterisk-monitor/app/ami"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Trunkstatus(ctx *gin.Context) {
	data, err := ami.TrunkStatus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
