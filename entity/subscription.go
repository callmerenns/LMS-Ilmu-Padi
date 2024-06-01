package entity

import (
	"time"

	"gorm.io/gorm"
)

// Initialize Struct Subscription
type Subscription struct {
	gorm.Model
	User_ID    string    `json:"user_id"`
	Start_Date time.Time `json:"start_date"`
	End_Date   time.Time `json:"end_date"`
	Status     bool      `json:"status"`
}
