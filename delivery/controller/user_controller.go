package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

// Initialize Struct User Controller
type UserController struct {
	userUsecase *usecase.UserUsecase
}

// Construction to Access User Controller
func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

// Get Profile By ID
func (ctrl *UserController) GetProfileByID(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	user, err := ctrl.userUsecase.GetProfileByID(userID)
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	common.SendSingleResponse(c, user, "Success")
}
