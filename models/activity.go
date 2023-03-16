package models

import "time"

type Activity struct {
	ActivityID int       `json:"activity_id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"type: varchar(255)"`
	Email      string    `json:"email" gorm:"type: varchar(255)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

}
