package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	Email              string            `json:"email,omitempty"`
	Password           string            `json:"password"`
	Role               string            `json:"role"`
	SubscriptionStatus string            `json:"subscription_status"`
	Courses            []Course          `json:"courses"`
	Subscriptions      Subscription      `json:"subscriptions"`
	VerificationToken  VerificationToken `json:"verification_token"`
	ResetToken         string
	ResetTokenExpiry   time.Time
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
