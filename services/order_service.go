package services

import (
	"food-ordering-system-project/models"
	"food-ordering-system-project/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
	MenuRepo  *repository.MenuRepository
}

func (s *OrderService) GetMenuByID(id string) (*models.Menu, error) {
	return s.MenuRepo.GetMenuByID(id)
}

func (s *OrderService) CreateOrder(
	customerName string,
	customerSurname string,
	phone string,
	address string,
	email string,
	menuID string,
	quantity string,
	status string,
	price string,
	total string,
) (int64, error) {

	return s.OrderRepo.CreateOrder(
		customerName,
		customerSurname,
		phone,
		address,
		email,
		menuID,
		quantity,
		status,
		price,
		total,
	)
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

func (s *OrderService) UpdateOrderStatus(
	id string,
	status string,
	userID int,
	reason string,
) error {

	return s.OrderRepo.UpdateOrderStatus(
		id,
		status,
		userID,
		reason,
	)
}

func (s *OrderService) GetOrderByID(id string) (models.Order, error) {
	return s.OrderRepo.GetOrderByID(id)
}

func (s *OrderService) GetOrderHistory(id string) ([]models.OrderHistory, error) {
	return s.OrderRepo.GetOrderHistory(id)
}

func (s *OrderService) CreateOrderHistory(
	orderID int64,
	status string,
) error {

	return s.OrderRepo.CreateOrderHistory(
		orderID,
		status,
	)
}