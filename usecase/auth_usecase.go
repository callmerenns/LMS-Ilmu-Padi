package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kelompok-2/ilmu-padi/client"
	"github.com/kelompok-2/ilmu-padi/config"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/service"
	"github.com/kelompok-2/ilmu-padi/shared/utils"
)

// Initialize Struct Auth Usecase
type AuthUsecase struct {
	authRepo    *repository.AuthRepository
	mailService *service.MailService
}

// Initialize Interface Email Sender Repository
type EmailSender interface {
	Send(to, subject, body string) error
}

// Construction to Access Auth Usecase
func NewAuthUsecase(authRepository *repository.AuthRepository, mailService *service.MailService) *AuthUsecase {
	return &AuthUsecase{authRepo: authRepository, mailService: mailService}
}

// Register
func (u *AuthUsecase) Register(data dto.RegisterDto) (*entity.User, error) {
	hashedPassword, err := config.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
		Role:     data.Role,
	}

	if err := u.authRepo.Create(user); err != nil {
		return nil, err
	}

	verificationToken, err := u.GenerateVerificationToken()
	if err != nil {
		return nil, err
	}

	user.VerificationToken = verificationToken
	user.VerificationTokenExpiry = time.Now().Add(24 * time.Hour)

	if err := u.authRepo.Update(user); err != nil {
		return nil, err
	}

	verificationURL := "http://ilmupadi.com/verify-email?token=" + verificationToken
	replacements := map[string]string{
		"{verificationLink}": verificationURL,
	}
	html := utils.FormatTemplate(client.VerifyEmailTemplate, replacements)

	if err := u.mailService.SendMail("Reset Password", html, []string{user.Email}); err != nil {
		return nil, err
	}

	return user, nil
}

// Login
func (u *AuthUsecase) Login(data dto.LoginDto) (string, error) {
	user, err := u.authRepo.FindByEmailAuth(data.Email)
	if err != nil {
		return "", err
	}

	if !config.CheckPasswordHash(data.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("YOUR_SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Logout
func (u *AuthUsecase) Logout() error {
	return nil
}

// Generate Reset Token
func (u *AuthUsecase) GenerateResetToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Forgot Password
func (u *AuthUsecase) ForgotPassword(data dto.ForgotPasswordDto) error {
	user, err := u.authRepo.FindByEmailAuth(data.Email)
	if err != nil {
		return errors.New("user not found")
	}

	token, err := u.GenerateResetToken()
	if err != nil {
		return err
	}

	user.ResetToken = token
	user.ResetTokenExpiry = time.Now().Add(1 * time.Hour)

	if err := u.authRepo.UpdateResetToken(user); err != nil {
		return err
	}

	resetURL := "http://ilmupadi.com/reset-password?token=" + token
	replacements := map[string]string{
		"{resetLink}": resetURL,
	}
	html := utils.FormatTemplate(client.ResetPasswordTemplate, replacements)

	return u.mailService.SendMail("Reset Password", html, []string{user.Email})
}

// Reset Password
func (u *AuthUsecase) ResetPassword(data dto.ResetPasswordDto) error {
	user, err := u.authRepo.FindByResetToken(data.Token)
	if err != nil {
		return errors.New("invalid token")
	}

	if time.Now().After(user.ResetTokenExpiry) {
		return errors.New("token expired")
	}

	hashedPassword, err := config.HashPassword(data.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.ResetToken = ""
	user.ResetTokenExpiry = time.Time{}

	return u.authRepo.Update(user)
}

// Generate Verification Token
func (u *AuthUsecase) GenerateVerificationToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Verify Email
func (u *AuthUsecase) VerifyEmail(token string) error {
	user, err := u.authRepo.FindByVerificationToken(token)
	if err != nil {
		return errors.New("invalid token")
	}

	if time.Now().After(user.VerificationTokenExpiry) {
		return errors.New("token expired")
	}

	user.Verified = true
	user.VerificationToken = ""
	user.VerificationTokenExpiry = time.Time{}

	return u.authRepo.Update(user)
}
