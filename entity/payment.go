package entity

import (
	"time"

	"gorm.io/gorm"
)

// Initialize Struct Payment
type Payment struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	User_ID        string    `json:"user_id"`
	Order_ID       string    `json:"order_id"`
	Transaction_ID string    `json:"transaction_id"`
	Amount         float64   `json:"amount"`
	Payment_Method string    `json:"payment_method"`
	Status         string    `json:"status"`
	Paid_At        time.Time `json:"paid_at"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
