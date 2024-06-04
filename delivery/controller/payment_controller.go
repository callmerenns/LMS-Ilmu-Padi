package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

// Initialize Struct Payment Controller
type PaymentController struct {
	paymentUsecase usecase.PaymentUsecase
	rg             *gin.RouterGroup
	authMid        middleware.AuthMiddleware
}

// Construction to Access Payment Controller
func NewPaymentController(paymentUsecase usecase.PaymentUsecase, rg *gin.RouterGroup, authMid middleware.AuthMiddleware) *PaymentController {
	return &PaymentController{paymentUsecase: paymentUsecase, rg: rg, authMid: authMid}
}

func (h *PaymentController) GetCampaignTransactions(c *gin.Context) {
	var input dto.GetCourseTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser

	transactions, err := h.paymentUsecase.GetTransactionsByCourseID(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, transactions)
}

func (h *PaymentController) CreateTransaction(c *gin.Context) {
	var input dto.CreateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		common.SendErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser

	newTransaction, err := h.paymentUsecase.CreateTransaction(input)

	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	common.SendSuccessResponse(c, http.StatusOK, newTransaction)
}

func (h *PaymentController) GetNotification(c *gin.Context) {
	var input dto.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.paymentUsecase.ProcessPayment(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)
}

// Routing Payment
func (pc *PaymentController) Route() {
	pc.rg.GET(routes.GetCourseTransaction, pc.authMid.RequireToken("admin", "instructor", "user"), pc.CreateTransaction)
	pc.rg.POST(routes.PostTransaction, pc.authMid.RequireToken("admin", "instructor", "user"), pc.CreateTransaction)
	pc.rg.POST(routes.GetNotification, pc.authMid.RequireToken("admin", "instructor", "user"), pc.CreateTransaction)
}
