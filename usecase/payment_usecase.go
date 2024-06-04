package usecase

import (
	"errors"
	"strconv"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/service"
)

type paymentUsecase struct {
	repository       repository.PaymentRepository
	courseRepository repository.CourseRepository
	paymentService   service.Service
}

type PaymentUsecase interface {
	GetTransactionsByCourseID(input dto.GetCourseTransactionsInput) ([]entity.Transaction, error)
	CreateTransaction(input dto.CreateTransactionInput) (entity.Transaction, error)
	ProcessPayment(input dto.TransactionNotificationInput) error
	GetAllTransactions() ([]entity.Transaction, error)
}

func NewPaymentUsecase(repository repository.PaymentRepository, courseRepository repository.CourseRepository, paymentService service.Service) *paymentUsecase {
	return &paymentUsecase{repository, courseRepository, paymentService}
}

func (s *paymentUsecase) GetTransactionsByCourseID(payload dto.GetCourseTransactionsInput) ([]entity.Transaction, error) {
	course, err := s.courseRepository.FindByID(payload.ID)
	if err != nil {
		return []entity.Transaction{}, err
	}

	if course.UserId != strconv.Itoa(payload.User.ID) {
		return []entity.Transaction{}, errors.New("not an owner of the course")
	}

	transactions, err := s.repository.GetByCourseID(payload.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *paymentUsecase) CreateTransaction(payload dto.CreateTransactionInput) (entity.Transaction, error) {
	transaction := entity.Transaction{}
	transaction.Course_ID = payload.CourseID
	transaction.Amount = payload.Amount
	transaction.User_ID = payload.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransacation := service.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransacation, payload.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.Payment_URL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *paymentUsecase) ProcessPayment(input dto.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	course, err := s.courseRepository.FindByID(updatedTransaction.Course_ID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		course.BackerCount = course.BackerCount + 1
		course.CurrentAmount = course.CurrentAmount + updatedTransaction.Amount

		_, err := s.courseRepository.Update(course)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *paymentUsecase) GetAllTransactions() ([]entity.Transaction, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
