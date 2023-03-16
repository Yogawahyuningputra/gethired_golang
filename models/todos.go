package models

import "time"

type Todos struct {
	TodoID          int       `json:"todo_id" gorm:"primaryKey"`
	Activity        Activity  `json:"activity" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ActivityGroupID"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title" gorm:"type: varchar(255)"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority" gorm:"type: varchar(255)"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
