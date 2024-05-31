package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id uint) (entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Save(user entity.User) error {
	return r.db.Save(&user).Error
}

func (r *UserRepository) Create(user entity.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) UpdateResetToken(user *entity.User) error {
	return r.db.Save(user).Error
}
