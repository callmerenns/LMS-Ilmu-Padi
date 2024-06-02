package controller

import (
	"net/http"
	"strconv"

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

// Find All
func (ctrl *UserController) GetList(c *gin.Context) {
	users, err := ctrl.userUsecase.FindAll()
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, users, "Success")
}

// Get Profile By ID
func (ctrl *UserController) GetProfileByID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}
	user, err := ctrl.userUsecase.GetProfileByID(uint(userID))
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	common.SendSingleResponse(c, user, "Success")
}
