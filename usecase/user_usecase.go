package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/kelompok-2/ilmu-padi/config"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/service"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository *repository.UserRepository
	emailSender    EmailSender // Interface for sending email
}

type EmailSender interface {
	Send(to, subject, body string) error
}

func NewUserUsecase(userRepository *repository.UserRepository, emailSender EmailSender) *UserUsecase {
	return &UserUsecase{userRepository: userRepository, emailSender: emailSender}
}

func (u *UserUsecase) Register(name, email, password string) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{Name: name, Email: email, Password: string(hashedPassword)}
	if err := u.userRepository.Create(user); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := service.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUsecase) GenerateResetToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (u *UserUsecase) ForgotPassword(email string) error {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	token, err := u.GenerateResetToken()
	if err != nil {
		return err
	}

	user.ResetToken = token
	user.ResetTokenExpiry = time.Now().Add(1 * time.Hour)

	if err := u.userRepository.UpdateResetToken(user); err != nil {
		return err
	}

	resetURL := "http://yourdomain.com/reset-password?token=" + token
	body := "To reset your password, please click the following link: " + resetURL

	return u.emailSender.Send(user.Email, "Password Reset Request", body)
}

func (u *UserUsecase) ResetPassword(token, newPassword string) error {
	user, err := u.userRepository.FindByResetToken(token)
	if err != nil {
		return errors.New("invalid token")
	}

	if time.Now().After(user.ResetTokenExpiry) {
		return errors.New("token expired")
	}

	hashedPassword, err := config.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.ResetToken = ""
	user.ResetTokenExpiry = time.Time{}

	return u.userRepository.Update(user)
}
