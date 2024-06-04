package dto

import (
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct Get Campaign Transactions Input
type GetCourseTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User entity.User
}

// Initialize Struct Create Transaction Input
type CreateTransactionInput struct {
	Amount   int `json:"amount" binding:"required"`
	CourseID int `json:"course_id" binding:"required"`
	User     entity.User
}

// Initialize Struct Transaction Notification Input
type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
