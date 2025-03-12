package event

import "github.com/google/uuid"

type UserMessage struct {
	BaseModelMessage
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	Role     string    `json:"role,omitempty"`
	IsActive bool      `json:"is_active,omitempty"`
}
