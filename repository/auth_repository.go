package repository

import (
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/service"
)

func Login(c *gin.Context) {
	var input entity.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := service.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}

func Register(c *gin.Context) {
	var input entity.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{Name: input.Name, Email: input.Email, Password: HashPassword(input.Password)}
	db.Create(&user)

	token := uuid.New().String()
	verificationToken := VerificationToken{UserID: user.ID, Token: token}
	db.Create(&verificationToken)

	link := "http://yourdomain.com/verify?token=" + token
	sendVerificationEmail(user.Email, link)

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful. Please check your email to verify your account."})
}

func VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	var verificationToken VerificationToken
	if err := db.Where("token = ?", token).First(&verificationToken).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	var user entity.User
	db.First(&user, verificationToken.UserID)
	user.IsVerified = true
	db.Save(&user)

	db.Delete(&verificationToken)

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func sendVerificationEmail(to string, link string) {
	from := "youremail@example.com"
	password := "yourpassword"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Verify your email\n\n" +
		"Please click the following link to verify your email:\n" + link

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", from, password, "smtp.example.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		panic(err)
	}
}
