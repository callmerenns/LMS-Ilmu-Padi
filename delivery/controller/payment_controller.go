package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// Initialize Struct Payment Controller
type PaymentController struct {
	paymentUsecase *usecase.PaymentUsecase
}

// Construction to Access Payment Controller
func NewPaymentController(paymentUsecase *usecase.PaymentUsecase) *PaymentController {
	return &PaymentController{paymentUsecase: paymentUsecase}
}

// Create Payment
func (ctrl *PaymentController) CreatePayment(c *gin.Context) {
	var input struct {
		OrderID     string `json:"order_id"`
		GrossAmount int64  `json:"gross_amount"`
		UserID      string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
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

	snapResp, err := ctrl.paymentUsecase.CreateSnapTransaction(snapReq)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	common.SendSingleResponse(c, snapResp.Token, "Success")
}
