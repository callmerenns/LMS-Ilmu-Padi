package service

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user entity.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user entity.User) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("missing env file %v", err.Error())
	}

	var (
		ClientKey string
		ServerKey string
	)

	ClientKey = os.Getenv("CLIENT_KEY")
	fmt.Println("Client Key : ", ClientKey)

	ServerKey = os.Getenv("SERVER_KEY")
	fmt.Println("Server Key : ", ServerKey)

	if ClientKey == "" || ServerKey == "" {
		return "", fmt.Errorf("client key or server key is not set")
	}

	midclient := midtrans.NewClient()
	midclient.ServerKey = ServerKey
	midclient.ClientKey = ClientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", fmt.Errorf("error getting snap token: %v", err)
	}

	return snapTokenResp.RedirectURL, nil
}
