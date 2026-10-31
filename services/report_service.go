package services

import (
	"food-ordering-system-project/models"
	"food-ordering-system-project/repository"
)

type ReportService struct {
	ReportRepo *repository.ReportRepository
}

func (s *ReportService) GetSalesReport(
	fromDate string,
	toDate string,
) (models.SalesReport, error) {

	return s.ReportRepo.GetSalesReport(
		fromDate,
		toDate,
	)
}

func (s *ReportService) GetMenuPerformanceReport() ([]models.MenuPerformance, error) {

	return s.ReportRepo.GetMenuPerformanceReport()
}

func (s *ReportService) GetCustomerOrdersReport() ([]models.CustomerOrderReport, error) {

	return s.ReportRepo.GetCustomerOrdersReport()
}

func (s *ReportService) GetOrdersReport(
	startDate string,
	endDate string,
	status string,
) ([]models.OrdersReport, error) {

	return s.ReportRepo.GetOrdersReport(
		startDate,
		endDate,
		status,
	)
}