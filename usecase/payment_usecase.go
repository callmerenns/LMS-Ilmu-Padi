package usecase

import (
	"time"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/midtrans/midtrans-go/snap"
)

// Initialize Struct Payment Usecase
type PaymentUsecase struct {
	snapClient        *snap.Client
	paymentRepository *repository.PaymentRepository
}

// Construction to Access Payment Usecase
func NewPaymentUsecase(snapClient *snap.Client, paymentRepository *repository.PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{
		snapClient:        snapClient,
		paymentRepository: paymentRepository,
	}
}

// Create Snap Transaction
func (u *PaymentUsecase) CreateSnapTransaction(req *snap.Request) (*snap.Response, error) {
	snapResp, err := u.snapClient.CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	payment := entity.Payment{
		User_ID:        req.CustomerDetail.FName,
		Order_ID:       req.TransactionDetails.OrderID,
		Transaction_ID: snapResp.Token,
		Amount:         float64(req.TransactionDetails.GrossAmt),
		Payment_Method: "midtrans",
		Status:         "pending",
		Paid_At:        time.Now(),
	}

	if err := u.paymentRepository.SavePayment(payment); err != nil {
		return nil, err
	}

	return snapResp, nil
}
