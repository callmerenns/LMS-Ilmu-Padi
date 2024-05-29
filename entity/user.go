package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 string         `json:"id"`
	Name               string         `json:"name"`
	Email              string         `json:"email"`
	Password           string         `json:"password"`
	Role               string         `json:"role"`
	SubscriptionStatus string         `json:"subscription_status"`
	Courses            []Course       `json:"courses"`
	Subscriptions      []Subscription `json:"subscriptions"`
	PaswordResets      []PaswordReset `json:"password_resets"`
}
