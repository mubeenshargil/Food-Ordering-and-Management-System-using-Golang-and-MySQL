package models

type CustomerOrderReport struct {
	CustomerName string
	Email        string
	TotalOrders  int
	TotalSpending float64
}