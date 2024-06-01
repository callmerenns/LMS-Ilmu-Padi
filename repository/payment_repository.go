package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct Payment Repository
type PaymentRepository struct {
	db *gorm.DB
}

// Construction to Access Payment Repository
func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

// Save Payment
func (r *PaymentRepository) SavePayment(payment entity.Payment) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in SavePayment")
	}

	return r.db.Create(&payment).Error
}
