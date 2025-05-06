package entities

import "time"

type AutoIdEntity struct {
	Id        int        `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Validate interface {
	IsValidate()
}
