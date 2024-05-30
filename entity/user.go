package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            string         `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name          string         `json:"name"`
	Email         string         `json:"email,omitempty"`
	Password      string         `json:"password"`
	Role          string         `json:"role"`
	Courses       []Course       `json:"courses"`
	Subscriptions []Subscription `json:"subscriptions"`
}
