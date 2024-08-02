package ami

import (
	"github.com/sirupsen/logrus"
	"github.com/tqcenglish/amigo-go/utils"
)

type EndpointList struct {
	Event       string `json:"event"`
	DeviceState string `json:"deviceState"`
	ObjectName  string `json:"objectName"`
	ObjectType  string `json:"objectType"`
}

func ExtnesionStatus() ([]*EndpointList, error) {
	var action = map[string]string{
		"Action": "PJSIPShowEndpoints",
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

	events := make([]*EndpointList, 0)
	for _, eventMap := range eventsMapArray {
		switch eventMap.Data["Event"] {
		// case "EndpointListComplete":
		case "EndpointList":
			event := &EndpointList{}
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
