package models

import "time"

type Condition struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
}

type ConditionStatus struct {
	ID            string    `json:"id"`
	Status        bool      `json:"status"`
	LastChangedBy string    `json:"last_changed_by"`
	LastChangedAt time.Time `json:"last_changed_at"`
}
