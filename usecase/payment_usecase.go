package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentUsecase struct {
	snapClient        *snap.Client
	paymentRepository *repository.PaymentRepository
}

func NewPaymentUsecase(snapClient *snap.Client, paymentRepository *repository.PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{
		snapClient:        snapClient,
		paymentRepository: paymentRepository,
	}
}

func (u *PaymentUsecase) CreateSnapTransaction(req *snap.Request) (*snap.Response, error) {
	snapResp, err := u.snapClient.CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	payment := entity.Payment{
		UserID:        req.CustomerDetail.FName,
		OrderID:       req.TransactionDetails.OrderID,
		Amount:        float64(req.TransactionDetails.GrossAmt),
		PaymentMethod: "midtrans",
		Status:        "pending",
		TransactionID: snapResp.TransactionID,
	}

	if err := u.paymentRepository.SavePayment(payment); err != nil {
		return nil, err
	}

	return snapResp, nil
}
