package usecase

import (
	"errors"
	"strconv"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/service"
)

// Initialize Struct User Payment Usecase
type paymentUsecase struct {
	repository       repository.PaymentRepository
	courseRepository repository.CourseRepository
	paymentService   service.Service
}

// Initialize Interface User Course Sender Usecase
type PaymentUsecase interface {
	GetTransactionsByCourseID(input dto.GetCourseTransactionsInput, user string) ([]entity.Transaction, error)
	CreateTransaction(input dto.CreateTransactionInput, user string) (entity.Transaction, error)
	ProcessPayment(input dto.TransactionNotificationInput, user string) error
	GetAllTransactions(user string) ([]entity.Transaction, error)
}

// Construction to Access Payment Courses Usecase
func NewPaymentUsecase(repository repository.PaymentRepository, courseRepository repository.CourseRepository, paymentService service.Service) *paymentUsecase {
	return &paymentUsecase{repository, courseRepository, paymentService}
}

// Get Transactions By Course ID
func (s *paymentUsecase) GetTransactionsByCourseID(payload dto.GetCourseTransactionsInput, user string) ([]entity.Transaction, error) {
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

// Create Transaction
func (s *paymentUsecase) CreateTransaction(payload dto.CreateTransactionInput, user string) (entity.Transaction, error) {
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

// Process Payment
func (s *paymentUsecase) ProcessPayment(input dto.TransactionNotificationInput, user string) error {
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

		err := s.courseRepository.Update(course)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get All Transactions
func (s *paymentUsecase) GetAllTransactions(user string) ([]entity.Transaction, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
