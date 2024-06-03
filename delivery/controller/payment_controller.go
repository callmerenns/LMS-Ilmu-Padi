package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
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

// Create Payment
func (pc *PaymentController) CreatePayment(c *gin.Context) {
	var input struct {
		OrderID     string `json:"order_id"`
		GrossAmount int64  `json:"gross_amount"`
		UserID      string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  input.OrderID,
			GrossAmt: input.GrossAmount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: input.UserID,
			Email: "customer@example.com",
		},
	}

	snapResp, err := pc.paymentUsecase.CreateSnapTransaction(snapReq)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, snapResp.Token, "Success")
}

// Routing Payment
func (pc *PaymentController) Route() {
	pc.rg.POST(routes.PostPayment, pc.authMid.RequireToken("admin", "instructor", "user"), pc.CreatePayment)
}
