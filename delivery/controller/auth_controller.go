package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

// Initialize Struct Auth Controller
type AuthController struct {
	authUsecase *usecase.AuthUsecase
}

// Construction to Access Auth Controller
func NewAuthController(authUsecase *usecase.AuthUsecase) *AuthController {
	return &AuthController{authUsecase: authUsecase}
}

// Register
func (ctrl *AuthController) Register(c *gin.Context) {
	input := dto.RegisterDto{}
	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	user, err := ctrl.authUsecase.Register(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	common.SendCreateResponse(c, user)
}

// Login
func (ctrl *AuthController) Login(c *gin.Context) {
	input := dto.LoginDto{}
	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	token, err := ctrl.authUsecase.Login(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusUnauthorized, "Status Unauthorized")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	common.SendSuccessResponse(c, http.StatusOK, "Login Success")
}

// Logout
func (ctrl *AuthController) Logout(c *gin.Context) {
	if err := ctrl.authUsecase.Logout(); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	// Clear the token cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
	})

	common.SendSuccessResponse(c, http.StatusOK, "Logout Success")
}

// Verify Email
func (ctrl *AuthController) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	err := ctrl.authUsecase.VerifyEmail(token)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, "Email Verified Success")
}

// Forgot Password
func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	input := dto.ForgotPasswordDto{}

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	if err := ctrl.authUsecase.ForgotPassword(input); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status internal Server Error")
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, "Password Reset Link Sent...")
}

// Reset Password
func (ctrl *AuthController) ResetPassword(c *gin.Context) {
	input := dto.ResetPasswordDto{}

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	if err := ctrl.authUsecase.ResetPassword(input); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, "Password Has Been Reset!")
}
