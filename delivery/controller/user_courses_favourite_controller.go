package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

type UserCoursesFavouriteController struct {
	userCoursesFavouriteUsecase usecase.IUserCoursesFavouriteUsecase
}

func NewUserCoursesFavouriteController(userCoursesFavouriteUsecase usecase.IUserCoursesFavouriteUsecase) *UserCoursesFavouriteController {
	return &UserCoursesFavouriteController{
		userCoursesFavouriteUsecase: userCoursesFavouriteUsecase,
	}
}

func (u *UserCoursesFavouriteController) AddOrRemoveCourseFavourite(c *gin.Context) {
	var ucf entity.UserCoursesFavourite

	if err := c.ShouldBindJSON(&ucf); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, msg := u.userCoursesFavouriteUsecase.AddOrRemoveToFavourite(ucf)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, msg)
}

func (u *UserCoursesFavouriteController) GetUserFavouriteList(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	favouriteList, err := u.userCoursesFavouriteUsecase.FindAllByUserID(uint(userID))
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, favouriteList, "Success")
}
