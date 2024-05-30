package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
)

type userUsecase struct {
	userRepo repository.IUserRepo
}

type IUserUsecase interface {
	FindByEmail(email string) (entity.User, error)
	Insert(payload entity.User) (string, error)
	UpdatePassword(email, password string) error
}

func (u *userUsecase) FindByEmail(email string) (entity.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userUsecase) Insert(payload entity.User) (string, error) {
	if err := u.userRepo.CheckDuplicateEmail(payload.Email); err != nil {
		return "", err
	}

	userID, err := u.userRepo.Insert(payload)

	return userID, err
}

func (u *userUsecase) UpdatePassword(email, password string) error {
	return u.userRepo.UpdatePassword(email, password)
}

func NewUserUsecase(userRepo repository.IUserRepo) IUserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
