package dto

import (
	"github.com/gin-gonic/gin"
)

func SendSingleResponse(c *gin.Context, data interface{}, code int, msg string) {

	c.JSON(code, SingleResponse{
		Status: Status{
			Code:    code,
			Message: msg,
			Error:   false,
		},
		Data: data,
	})
}

func SendManyResponse(c *gin.Context, data []interface{}, paging Paging, code int, msg string) {

	c.JSON(code, ManyResponse{
		Status: Status{
			Code:    code,
			Message: msg,
			Error:   false,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendErrorResponse(c *gin.Context, code int, msg string, err any) {
	c.AbortWithStatusJSON(code, SingleResponse{
		Status: Status{
			Code:    code,
			Error:   err,
			Message: msg,
		},
		Data: nil,
	})
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

type Paging struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalRows  int `json:"totalRows"`
	TotalPages int `json:"totalPages"`
}

type ManyResponse struct {
	Status Status        `json:"status"`
	Data   []interface{} `json:"data"`
	Paging Paging        `json:"paging"`
}

type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}
