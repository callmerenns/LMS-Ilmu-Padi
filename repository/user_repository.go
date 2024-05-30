package repository

import (
	"errors"

	"github.com/kelompok-2/ilmu-padi/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepo struct {
	courseRepo       ICourseRepo
	subscribtionRepo ISubscribtionRepo
	db               *gorm.DB
}

func (n *userRepo) UpdatePassword(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = n.db.Model(&entity.User{}).Update("password", string(hashedPassword)).Where("email = ?", email).Error
	if err != nil {
		return err
	}

	return nil
}

func (n *userRepo) CheckDuplicateEmail(email string) error {
	var count int64
	err := n.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email already exists in database")
	}

	return nil
}

func (n *userRepo) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := n.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}

	user.Courses, err = n.courseRepo.FindUserCoursesByUserID(user.ID)
	if err != nil {
		return entity.User{}, err
	}

	user.Subscriptions, err = n.subscribtionRepo.FindUserSubscribtionsByUserID(user.ID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (n *userRepo) Insert(payload entity.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	payload.Password = string(hashedPassword)

	err = n.db.Create(&payload).Error

	if err != nil {
		return "", err
	}

	return payload.ID, nil
}

type IUserRepo interface {
	UpdatePassword(email, password string) error
	FindByEmail(email string) (entity.User, error)
	CheckDuplicateEmail(email string) error
	Insert(payload entity.User) (string, error)
}

func NewUserRepo(db *gorm.DB, courseRepo ICourseRepo, subscribtionRepo ISubscribtionRepo) IUserRepo {
	return &userRepo{
		db:               db,
		subscribtionRepo: subscribtionRepo,
		courseRepo:       courseRepo,
	}
}
