package main

import (
	"asterisk-monitor/api"
	"asterisk-monitor/app/ami"
	"asterisk-monitor/app/mysql"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once

func main() {
	go ami.StartAMI(func() {
		logrus.Info("ami callback function")
		// 首次连接才进行初始化
		once.Do(func() {
			// 周期性查询队列信息
			// IntervalTime := 2 * time.Second        // 觸發間隔時間
			// ticker := time.NewTicker(IntervalTime) // 設定 2 秒觸發一次
			// for {
			// 	c := <-ticker.C
			// 	fmt.Println("now: ", c)
			// 	ami.QueueStatus()
			// }

		})
	}, []func(event map[string]string){})

	go mysql.CreateDBInstance()
	api.HttpServer()
}
