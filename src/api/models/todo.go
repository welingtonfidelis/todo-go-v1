package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status" validate:"required,oneof=done doing todo"`
	Date        time.Time `json:"date" validate:"required"`
}

type CreateTodoInput struct {
	Name        string    `json:"name" validate:"required"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status" validate:"required,oneof=done doing todo"`
	Date        time.Time `json:"date" validate:"required"`
}
