package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

// Example function to save payment details
func (r *PaymentRepository) SavePayment(payment entity.Payment) error {
	return r.db.Create(&payment).Error
}
