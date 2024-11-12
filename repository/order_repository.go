package repository

import (
	"database/sql"
	"log"

	"github.com/book-online-shop/logger"
	"github.com/book-online-shop/model"
	"go.uber.org/zap"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return OrderRepository{db: db}
}

func (r *OrderRepository) Create(payload *model.Order) error {
	query := `INSERT INTO orders(name, quantity, diskon, amount, final_amount) VALUES($1, $2, $3, $4, $5)`
	logger.Logger.Info("Executing query",
		zap.String("repository", "OrderRepository"),
		zap.String("query", query),
		zap.String("name", payload.Name),
		zap.Int("quantity", payload.Quantity),
		zap.Float64("amount", payload.Amount),
		zap.Float64("final_amount", payload.FinalAmount),
	)
	_, err := r.db.Exec(query, payload.Name, payload.Quantity, payload.Diskon, payload.Amount, payload.FinalAmount)
	log.Printf("Amount: %f, Discount: %f, FinalAmount: %f", payload.Amount, payload.Diskon, payload.FinalAmount)

	if err != nil {
		logger.Logger.Error("Failed to execute query",
			zap.String("repository", "OrderRepository"),
			zap.String("query", query),
			zap.Error(err),
		)
		log.Printf("Error executing query: %v\n", err)
		return err
	}
	return nil
}
