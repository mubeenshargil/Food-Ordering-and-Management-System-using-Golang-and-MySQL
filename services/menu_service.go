package services

import (
	"errors"
	"strconv"
	"food-ordering-system-project/models"
	"food-ordering-system-project/repository"
)

type MenuService struct {
	MenuRepo *repository.MenuRepository
}

func (s *MenuService) GetMenus() ([]models.Menu, error) {

	return s.MenuRepo.GetMenus()
}

func (s *MenuService) GetMenuByID(
	id string,
) (*models.Menu, error) {

	return s.MenuRepo.GetMenuByID(id)
}

func (s *MenuService) CreateMenu(
	name string,
	description string,
	price string,
	image string,
	status string,
) error {
     
	if name == "" {
		return errors.New("menu name is required")
	}

	if description == "" {
		return errors.New("menu description is required")
	}

	if image == "" {
		return errors.New("menu image is required")
	}

	if status == "" {
		return errors.New("menu status is required")
	}

	priceValue, err := strconv.ParseFloat(price, 64)

	if err != nil {
		return errors.New("invalid price")
	}

	if priceValue <= 0 {
		return errors.New("price must be greater than 0")
	}

	return s.MenuRepo.CreateMenu(
		name,
		description,
		price,
		image,
		status,
	)
}

func (s *MenuService) UpdateMenu(
	id string,
	name string,
	description string,
	price string,
	status string,
) error {

	if id == "" {
		return errors.New("menu id is required")
	}

	if name == "" {
		return errors.New("menu name is required")
	}

	if description == "" {
		return errors.New("menu description is required")
	}

	if status == "" {
		return errors.New("menu status is required")
	}

	priceValue, err := strconv.ParseFloat(price, 64)

	if err != nil {
		return errors.New("invalid price")
	}

	if priceValue <= 0 {
		return errors.New("price must be greater than 0")
	}


	return s.MenuRepo.UpdateMenu(
		id,
		name,
		description,
		price,
		status,
	)
}