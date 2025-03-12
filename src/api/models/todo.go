package models

import "time"

type Todo struct {
	Name        string    `json:"name" validate:"required"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status" validate:"required,oneof=done doing to-do"`
	Date        time.Time `json:"date" validate:"required"`
}
