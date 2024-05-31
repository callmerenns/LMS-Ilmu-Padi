package entity

import (
	"time"

	"gorm.io/gorm"
)

type VerificationToken struct {
	gorm.Model
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
