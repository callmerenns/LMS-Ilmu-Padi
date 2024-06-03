package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct Payment Repository
type paymentRepository struct {
	db *gorm.DB
}

// Initialize Interface Payment Sender Repository
type PaymentRepository interface {
	SavePayment(payment entity.Payment) error
}

// Construction to Access Payment Repository
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// Save Payment
func (p *paymentRepository) SavePayment(payment entity.Payment) error {
	if p.db == nil {
		log.Fatal("Database connection is nil in SavePayment")
	}

	return p.db.Create(&payment).Error
}
