package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
)

// Initialize Struct User Usecase
type UserUsecase struct {
	userRepo *repository.UserRepository
}

// Construction to Access User Usecase
func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// Get Profile By ID
func (u *UserUsecase) GetProfileByID(userID uint) (*entity.User, error) {
	return u.userRepo.FindByID(userID)
}
