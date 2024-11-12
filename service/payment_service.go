package service

import (
	"fmt"

	"github.com/book-online-shop/model"
	"github.com/book-online-shop/repository"
	// "github.com/google/uuid"
)

type PaymentService struct {
	paymentRepo *repository.PaymentRepository
}

func NewPaymentService(paymentRepo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *PaymentService) CreatePayment(payload model.Payment) error {
	// if payload.ID == "" || payload.Name == "" || payload.Photo == true {
	// 	return fmt.Errorf("all payload is required")
	// }
	// payload.ID = uuid.New().String()

	err := s.paymentRepo.Create(&payload)
	if err != nil {
		return fmt.Errorf("failed to create payment: %s", err)
	}
	return nil
}
