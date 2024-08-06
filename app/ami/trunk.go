package ami

import (
	"github.com/sirupsen/logrus"
	"github.com/tqcenglish/amigo-go/utils"
)

type OutboundRegistrationDetail struct {
	Event      string `json:"event"`
	Endpoint   string `json:"endpoint"`
	ObjectName string `json:"objectName"`
	Status     string `json:"status"`
}

// 获取中继状态
func TrunkStatus() ([]*OutboundRegistrationDetail, error) {
	var action = map[string]string{
		"Action": "PJSIPShowRegistrationsOutbound",
	}
	responseMap, eventsMapArray, err := AminInstance.Send(action)

	if err != nil {
		return nil, err
	}

	response := &AMIRes{}
	for k, v := range responseMap {
		if err := utils.SetField(response, k, v); err != nil {
			logrus.Errorf("Response SetField error %+v", err)
		}
	}

	events := make([]*OutboundRegistrationDetail, 0)
	for _, eventMap := range eventsMapArray {
		switch eventMap.Data["Event"] {
		// case "ContactStatusDetailComplete":
		case "OutboundRegistrationDetail":
			// logrus.Info(eventMap.Data)
			event := &OutboundRegistrationDetail{}
			for k, v := range eventMap.Data {
				if err := utils.SetField(event, k, v); err != nil {
					logrus.Errorf("Event SetField error %+v", err)
					continue
				}
			}
			events = append(events, event)
		}

	}
	return events, nil
}
