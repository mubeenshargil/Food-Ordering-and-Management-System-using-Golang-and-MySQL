package models

type OrdersReport struct {
	OrderID      int
	CustomerName string
	Email        string
	Status       string
	Reason       string
	Total        float64
	OrderTime    string
}