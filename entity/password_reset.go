package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaswordReset struct {
	gorm.Model
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}
