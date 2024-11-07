package service

import (
	"fmt"

	"github.com/book-online-shop/model"
	"github.com/book-online-shop/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterNewUser(payload model.User) error {
	if payload.ID == "" || payload.Username == "" || payload.Password == "" {
		return fmt.Errorf("all payload is required")
	}

	err := s.userRepo.Create(&payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}
	return nil
	// userRepo := repository.NewUserRepository(&payload)
}

func (s *UserService) LoginService(payload *model.User) error {
	err := s.userRepo.Login(payload)
	if err != nil {
		return err
	}

	return nil
}
