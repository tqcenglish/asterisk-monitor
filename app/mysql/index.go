package mysql

import (
	"fmt"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var DBOrmInstance *xorm.Engine

// TODO 不使用 mysql， 代码暂时保留
func CreateDBInstance() {
	var err error
	// DBOrmInstance, err = xorm.NewEngine("sqlite3", "playcall.db")
	url := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8",
		"root",
		"123456",
		"192.168.18.252",
		"ippbx_db",
	)
	// logrus.Infof("mysql url %s", url)
	DBOrmInstance, err = xorm.NewEngine("mysql", url)
	if err != nil {
		logrus.Panic(err)
		return
	}
	err = DBOrmInstance.Ping()
	if err != nil {
		logrus.Error("mysql ping error: ", err)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		return
	}
	DBOrmInstance.ShowSQL(true)
	// info, err := os.Open(configs.ConfigGlobal.LogInfoPath)
	// if err != nil {
	// 	logrus.Error(err.Error())
	// 	return
	// }
	// DBOrmInstance.SetLogger(log.NewSimpleLogger(info))
	DBOrmInstance.Dialect().URI().Timeout = 30 * time.Second
}
