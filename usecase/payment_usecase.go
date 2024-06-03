package usecase

import (
	"time"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/midtrans/midtrans-go/snap"
)

// Initialize Struct Payment Usecase
type paymentUsecase struct {
	snapClient        snap.Client
	paymentRepository repository.PaymentRepository
}

// Initialize Interface Payment Sender Usecase
type PaymentUsecase interface {
	CreateSnapTransaction(req *snap.Request) (*snap.Response, error)
}

// Construction to Access Payment Usecase
func NewPaymentUsecase(snapClient snap.Client, paymentRepository repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{
		snapClient:        snapClient,
		paymentRepository: paymentRepository,
	}
}

// Create Snap Transaction
func (p *paymentUsecase) CreateSnapTransaction(req *snap.Request) (*snap.Response, error) {
	snapResp, err := p.snapClient.CreateTransaction(req)
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

	if err := p.paymentRepository.SavePayment(payment); err != nil {
		return nil, err
	}

	return snapResp, nil
}
