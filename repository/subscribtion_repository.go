package repository

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"gorm.io/gorm"
)

type subscribtionRepo struct {
	db *gorm.DB
}

func (s *subscribtionRepo) FindUserSubscribtionsByUserID(uid string) ([]entity.Subscription, error) {
	var subscriptions []entity.Subscription
	err := s.db.Where("user_id = ?", uid).Find(&subscriptions).Error
	return subscriptions, err
}

type ISubscribtionRepo interface {
	FindUserSubscribtionsByUserID(uid string) ([]entity.Subscription, error)
}

func NewSubscribtionRepo(db *gorm.DB) ISubscribtionRepo {
	return &subscribtionRepo{
		db: db,
	}
}
