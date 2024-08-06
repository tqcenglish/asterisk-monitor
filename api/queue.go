package api

import (
	"asterisk-monitor/app/ami"
	"asterisk-monitor/app/mysql"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var QueueNumberName map[string]string
var once sync.Once

type QueueModel struct {
	Name  string `xorm:"'name'" json:"name"`
	Alias string `xorm:"'alias'" json:"alias"`
}

func (*QueueModel) TableName() string {
	return "t_queue"
}

func InitQueueInfo() {
	var models []QueueModel
	err := mysql.DBOrmInstance.Find(&models)
	if err != nil {
		logrus.Error(err)
		return
	}
	QueueNumberName = make(map[string]string)
	for _, data := range models {
		QueueNumberName[data.Name] = data.Alias
	}
	logrus.Infof("%+v", QueueNumberName)
}
func QueueStatus(ctx *gin.Context) {
	once.Do(InitQueueInfo)

	data1, _, err := ami.QueueStatus()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resut := make([]*ami.QueueParamsEvent, 0)
	for _, event := range data1 {
		resut = append(resut, &ami.QueueParamsEvent{
			Event:     event.Event,
			Calls:     event.Calls,
			Abandoned: event.Abandoned,
			Completed: event.Completed,
			Queue:     event.Queue,
			Name:      QueueNumberName[event.Queue],
		})
	}
	ctx.JSON(http.StatusOK, resut)
}
