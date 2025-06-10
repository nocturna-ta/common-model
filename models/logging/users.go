package logging

import "github.com/nocturna-ta/common-model/models/event"

type KPULogs struct {
	event.BaseModelMessage
	Name     string `json:"name,omitempty"`
	KPUName  string `json:"kpu_name,omitempty"`
	KPUType  string `json:"kpu_type,omitempty"`
	Activity string `json:"activity,omitempty"`
}
