package todosdto

import "time"

type TodosRequest struct {
	ActivityGroupID int    `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

type TodosUpdate struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
}

type TodosResponse struct {
	TodoID          int       `json:"todo_id"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
