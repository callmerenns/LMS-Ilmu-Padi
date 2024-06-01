package entity

import (
	"time"

	"gorm.io/gorm"
)

// Initialize Struct User
type User struct {
	gorm.Model
	Name                    string    `json:"name"`
	Email                   string    `json:"email,omitempty"`
	Password                string    `json:"password"`
	Role                    string    `json:"role"`
	Subscription_Status     string    `json:"subscription_status"`
	VerificationToken       string    `json:"verification_token"`
	VerificationTokenExpiry time.Time `json:"verification_token_expiry"`
	ResetToken              string    `json:"reset_token"`
	ResetTokenExpiry        time.Time `json:"reset_token_expiry"`
	Verified                bool      `json:"verified"`
}
