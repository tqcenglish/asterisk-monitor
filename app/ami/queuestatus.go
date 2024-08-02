package ami

import (
	"github.com/sirupsen/logrus"
	"github.com/tqcenglish/amigo-go/utils"
)

type AMIRes struct {
	Response  string `json:"response"`
	ActionID  string `json:"action_id"`
	EventList string `json:"event_list"`
	Message   string `json:"message"`
}

// Event: QueueMember
// Queue: 0310
// Name: 856
// Location: Local/856@app-push
// StateInterface: Local/856@app-push
// Membership: static
// Penalty: 0
// CallsTaken: 0
// LastCall: 0
// LastPause: 0
// LoginTime: 1722503083
// InCall: 0
// Status: 1
// Paused: 0
// PausedReason:
// Wrapuptime: 0

type QueueMemberEvent struct {
	Event          string `json:"event"`
	Queue          string `json:"queue"`
	Name           string `json:"name"`
	Location       string `json:"location"`
	StateInterface string `json:"state_interface"`
	Membership     string `json:"membership"`
	Penalty        string `json:"penalty"`
	CallsTaken     string `json:"calls_taken"`
	LastCall       string `json:"last_call"`
	LastPause      string `json:"last_pause"`
	LoginTime      string `json:"login_time"`
	InCall         string `json:"in_call"`
	Status         string `json:"status"`
	Paused         string `json:"paused"`
	PausedReason   string `json:"paused_reason"`
	Wrapuptime     string `json:"wrapuptime"`
}

// Event: QueueParams
// Queue: 0306
// Max: 0
// Strategy: ringall
// Calls: 0
// Holdtime: 0
// TalkTime: 0
// Completed: 0
// Abandoned: 0
// ServiceLevel: 0
// ServicelevelPerf: 0.0
// ServicelevelPerf2: 0.0
// Weight: 0

type QueueParamsEvent struct {
	Event             string `json:"event"`
	Queue             string `json:"queue"`
	Max               string `json:"max"`
	Strategy          string `json:"strategy"`
	Calls             string `json:"calls"`
	Holdtime          string `json:"holdtime"`
	TalkTime          string `json:"talktime"`
	Completed         string `json:"completed"`
	Abandoned         string `json:"abandoned"`
	ServiceLevel      string `json:"servicelevel"`
	ServicelevelPerf  string `json:"servicelevel_perf"`
	ServicelevelPerf2 string `json:"servicelevel_perf2"`
	Weight            string `json:"weight"`
}

func QueueStatus() ([]*QueueParamsEvent, []*QueueMemberEvent, error) {
	var action = map[string]string{
		"Action": "QueueStatus",
	}
	responseMap, eventsMapArray, err := AminInstance.Send(action)

	if err != nil {
		return nil, nil, err
	}

	response := &AMIRes{}
	for k, v := range responseMap {
		if err := utils.SetField(response, k, v); err != nil {
			logrus.Errorf("Response SetField error %+v", err)
		}
	}
	events1 := make([]*QueueParamsEvent, 0)
	events2 := make([]*QueueMemberEvent, 0)
	for _, eventMap := range eventsMapArray {
		switch eventMap.Data["Event"] {
		case "QueueParams":
			event := &QueueParamsEvent{}
			for k, v := range eventMap.Data {
				if err := utils.SetField(event, k, v); err != nil {
					logrus.Errorf("Event SetField error %+v", err)
					continue
				}
			}
			events1 = append(events1, event)
		case "QueueMember":
			event := &QueueMemberEvent{}
			for k, v := range eventMap.Data {
				if err := utils.SetField(event, k, v); err != nil {
					logrus.Errorf("Event SetField error %+v", err)
					continue
				}
			}
			events2 = append(events2, event)
		}

	}

	return events1, events2, nil
}
