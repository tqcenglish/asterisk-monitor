package ami

import (
	"github.com/sirupsen/logrus"
	"github.com/tqcenglish/amigo-go"
	"github.com/tqcenglish/amigo-go/pkg"
)

var AminInstance *amigo.Amigo

func StartAMI(connectOKCallBack func(), handleEvents []func(event map[string]string)) {
	logrus.Info("Start AMI")
	settings := &amigo.Settings{
		Host:     "192.168.18.252",
		Port:     "5038",
		Username: "admin",
		Password: "admin",
		LogLevel: logrus.InfoLevel}
	logrus.Infof("ami setting: %+v", settings)
	AminInstance = amigo.New(settings, nil)
	AminInstance.EventOn(func(payload ...interface{}) {
		event := payload[0].(map[string]string)
		handleAMI(event)

	})
	AminInstance.ConnectOn(func(payload ...interface{}) {
		if payload[0] == pkg.Connect_OK {
			logrus.Info("ami connect ok")
			connectOKCallBack()
		} else {
			logrus.Errorf("ami connect failure %+v", payload)
		}
	})
	AminInstance.Connect()
}

func handleAMI(event map[string]string) {
	logrus.Debug("ami event: ", event)
	switch event["Event"] {
	case "ContactStatus":
	}
}

func Connected() bool {
	if AminInstance != nil {
		return AminInstance.Connected()
	}
	return false
}
