package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/common/dto"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

type userCtrl struct {
	userUsecase usecase.IUserUsecase
}

func (ct *userCtrl) CreateCtrl(c *gin.Context) {
	var payload entity.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	_, err := ct.userUsecase.Insert(payload)
	if err != nil {
		dto.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	dto.SendSingleResponse(c, nil, http.StatusCreated, "Created")
}

func (ct *userCtrl) Routing(r *gin.RouterGroup) {
	r.POST("/users", ct.CreateCtrl)
}

func NewUserCtrl(userUsecase usecase.IUserUsecase) *userCtrl {
	return &userCtrl{
		userUsecase: userUsecase,
	}
}
