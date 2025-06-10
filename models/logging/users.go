package logging

import "github.com/nocturna-ta/common-model/models/event"

type KPULogs struct {
	event.BaseModelMessage
	UserID   string `json:"user_id,omitempty"`
	Address  string `json:"address,omitempty"`
	Role     string `json:"role,omitempty"`
	Activity string `json:"activity,omitempty"`
}
