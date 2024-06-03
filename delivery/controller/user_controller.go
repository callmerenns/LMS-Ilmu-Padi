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

<<<<<<< HEAD
// Construction to Access User Controller
func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}
=======
func (ct *userCtrl) CreateCtrl(c *gin.Context) {
	var payload entity.User
>>>>>>> dev/tsaqif

// Find All
func (ctrl *UserController) GetList(c *gin.Context) {
	users, err := ctrl.userUsecase.FindAll()
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, users, "Success")
}

<<<<<<< HEAD
// Get Profile By ID
func (ctrl *UserController) GetProfileByID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
=======
func (ct *userCtrl) Routing(r *gin.RouterGroup) {
	r.POST("/users", ct.CreateCtrl)
}

func NewUserCtrl(userUsecase usecase.IUserUsecase) *userCtrl {
	return &userCtrl{
		userUsecase: userUsecase,
>>>>>>> dev/tsaqif
	}
	user, err := ctrl.userUsecase.GetProfileByID(uint(userID))
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	common.SendSingleResponse(c, user, "Success")
}
