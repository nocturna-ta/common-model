package event

import "time"

type BaseModelMessage struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsDeleted bool      `json:"is_deleted,omitempty"`
}
