package services

import (
	"errors"
	"strings"
	"food-ordering-system-project/models"
	"food-ordering-system-project/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) GetUserByEmail(email string) (*repository.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *UserService) CreateUser(
	name string,
	surname string,
	phone string,
	email string,
	password string,
	plainPassword string,
	role string,
	status string,
) error {
    
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(surname) == "" {
		return errors.New("surname is required")
	}

	if strings.TrimSpace(phone) == "" {
		return errors.New("phone is required")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(email, "@") {
		return errors.New("invalid email address")
	}

	if len(plainPassword) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if role != "admin" &&
		role != "cashier" &&
		role != "customer" {

		return errors.New("invalid role")
	}

	if status == "" {
		return errors.New("status is required")
	}


	return s.UserRepo.CreateUser(
		name,
		surname,
		phone,
		email,
		password,
		plainPassword,
		role,
		status,
	)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) UpdateUser(
	id string,
	name string,
	surname string,
	phone string,
	email string,
	role string,
	status string,
) error {

	if id == "" {
		return errors.New("user id is required")
	}

	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(surname) == "" {
		return errors.New("surname is required")
	}

	if strings.TrimSpace(phone) == "" {
		return errors.New("phone is required")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(email, "@") {
		return errors.New("invalid email address")
	}

	if role != "admin" &&
		role != "cashier" &&
		role != "customer" {

		return errors.New("invalid role")
	}

	if status == "" {
		return errors.New("status is required")
	}

	return s.UserRepo.UpdateUser(
		id,
		name,
		surname,
		phone,
		email,
		role,
		status,
	)
}