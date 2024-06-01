package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct User Repository
type UserRepository struct {
	db *gorm.DB
}

// Construction to Access User Repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	if db == nil {
		log.Fatal("Database connection is nil UserRepository")
	}

	return &UserRepository{db: db}
}

// Find By ID
func (r *UserRepository) FindByID(id uint) (*entity.User, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in FindByID")
	}

	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Get Roles By User ID
func (r *UserRepository) GetRolesByUserID(userID uint) ([]entity.User, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in GetRolesByUserID")
	}

	var roles []entity.User
	if err := r.db.Joins("JOIN user_roles on roles.id = user_roles.role_id").Where("user_roles.user_id = ?", userID).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// Find By Email User
func (r *UserRepository) FindByEmailUser(email string) (*entity.User, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in FindByEmailUser")
	}
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
