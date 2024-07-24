package main

import (
	"asterisk-monitor/api"
	"asterisk-monitor/app/ami"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once

func main() {
	go ami.StartAMI(func() {
		logrus.Info("ami callback function")
		// 首次连接才进行初始化
		once.Do(func() {
			// initAsterisk()
			// agi.InitPageingPlan(configs.ConfigGlobal.AsteriskPagingPath)
		})
	}, []func(event map[string]string){})

	api.HttpServer()
}
