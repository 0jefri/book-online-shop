package repository

import (
	"database/sql"

	"github.com/book-online-shop/model"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(payload *model.Payment) error {
	query := `INSERT INTO payment(name, photo, is_active) VALUES($1, $2, $3)`
	err := r.db.QueryRow(query, payload.Name, payload.Photo, payload.IsActive).Scan(&payload.Name, &payload.Photo, &payload.IsActive)
	if err != nil {
		return err
	}
	return nil
}
