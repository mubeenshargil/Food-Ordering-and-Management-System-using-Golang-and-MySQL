package services

import (
	"food-ordering-system-project/models"
	"food-ordering-system-project/repository"
)

type AdminService struct {
	MenuRepo *repository.MenuRepository
}

func (s *AdminService) GetMenus() ([]models.Menu, error) {

	return s.MenuRepo.GetMenus()
}

func (s *AdminService) GetDashboardStats() (*models.DashboardStats, error) {

	return repository.GetDashboardStats()
}