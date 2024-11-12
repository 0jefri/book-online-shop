package service

import (
	"log"

	"github.com/book-online-shop/logger"
	"github.com/book-online-shop/model"
	"github.com/book-online-shop/repository"
	"go.uber.org/zap"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(payload *model.Order) error {

	logger.Logger.Info("Processing order",
		zap.String("name", payload.Name),
		zap.Int("quantity", payload.Quantity),
		zap.Float64("amount", payload.Amount),
	)

	discountAmount := payload.Amount * (payload.Diskon / 100)
	finalAmount := payload.Amount - discountAmount

	payload.FinalAmount = finalAmount

	log.Printf("Amount: %f, Discount: %f, FinalAmount: %f", payload.Amount, payload.Diskon, finalAmount)

	err := s.repo.Create(payload)
	if err != nil {
		logger.Logger.Error("Failed to create order in repository",
			zap.String("repository", "OrderRepository"),
			zap.String("query", "INSERT INTO orders(name, quantity, diskon, amount, final_amount) VALUES($1, $2, $3, $4, $5)"),
			zap.Error(err),
		)
		return err
	}
	logger.Logger.Info("Order created successfully",
		zap.String("order_id", payload.ID),
	)

	return nil
}
