package repository

import (
	"database/sql"

	"github.com/book-online-shop/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) Create(payload *model.User) error {
	query := `INSERT INTO users(username, password) VALUES($1, $2)`
	_, err := r.db.Exec(query, payload.Username, payload.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Login(payload *model.User) error {
	query := `SELECT username, password users WHERE username=$1 AND password=$2`
	err := r.db.QueryRow(query, payload.Username, payload.Password).Scan(&payload.Username, &payload.Password)
	if err != nil {
		return err
	}
	return nil
}
