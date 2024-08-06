package api

import (
	"asterisk-monitor/app/mysql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 统计通话记录

type CallLogModel struct {
	Id         int    `xorm:"'id'" json:"id"`
	StartTime  string `xorm:"'start_time'" json:"startTime"`
	CallStatus string `xorm:"'call_status'" json:"callStatus"`
	CallType   string `xorm:"'call_type'" json:"callType"`
	Trunk      string `xorm:"'trunk'" json:"trunk"`
}

func (*CallLogModel) TableName() string {
	return "t_call_log"
}

func CallLog(ctx *gin.Context) {
	internalCallCount, err := mysql.DBOrmInstance.Where("start_time > ?", time.Now().Format("2006-01-02")).Count(&CallLogModel{CallType: "internal"})
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	incomingCallCount, err := mysql.DBOrmInstance.Where("start_time > ?", time.Now().Format("2006-01-02")).Count(&CallLogModel{CallType: "incoming"})
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	outgoingCallCount, err := mysql.DBOrmInstance.Where("start_time > ?", time.Now().Format("2006-01-02")).Count(&CallLogModel{CallType: "outgoing"})
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"internalCallCount": internalCallCount,
		"incomingCallCount": incomingCallCount,
		"outgoingCallCount": outgoingCallCount,
	})
}
