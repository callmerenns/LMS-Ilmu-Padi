package entity

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    bool      `json:"status"`
}
