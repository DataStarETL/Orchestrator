package models

import "time"

type Condition struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
}
