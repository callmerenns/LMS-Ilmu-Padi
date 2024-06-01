package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct Auth Repository
type AuthRepository struct {
	db *gorm.DB
}

// Construction to Access Auth Repository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	if db == nil {
		log.Fatal("Database connection is nil AuthRepository")
	}

	return &AuthRepository{db: db}
}

// Find By Reset Token
func (r *AuthRepository) FindByResetToken(token string) (*entity.User, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in FindByResetToken")
	}

	var user entity.User
	if err := r.db.Where("reset_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find by Verification Token
func (r *AuthRepository) FindByVerificationToken(token string) (*entity.User, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in FindByVerificationToken")
	}

	var user entity.User
	if err := r.db.Where("verification_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find By Email Auth
func (r *AuthRepository) FindByEmailAuth(email string) (*entity.User, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in FindByEmail")
	}

	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Save
func (r *AuthRepository) Save(user entity.User) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in Save")
	}

	return r.db.Save(&user).Error
}

// Create
func (r *AuthRepository) Create(user *entity.User) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in Create")
	}

	return r.db.Create(&user).Error
}

// Update Reset Token
func (r *AuthRepository) UpdateResetToken(user *entity.User) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in UpdateResetToken")
	}

	return r.db.Save(user).Error
}

// Update
func (r *AuthRepository) Update(user *entity.User) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in Update")
	}

	return r.db.Save(user).Error
}
