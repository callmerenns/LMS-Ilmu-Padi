package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/usecase"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentController struct {
	paymentUsecase *usecase.PaymentUsecase
}

func NewPaymentController(paymentUsecase *usecase.PaymentUsecase) *PaymentController {
	return &PaymentController{paymentUsecase: paymentUsecase}
}

func (ctrl *PaymentController) CreatePayment(c *gin.Context) {
	var input struct {
		OrderID     string `json:"order_id"`
		GrossAmount int64  `json:"gross_amount"`
		UserID      string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"snap_token": snapResp.Token})
}
