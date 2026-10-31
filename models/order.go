package models

type Order struct {
	ID        int
	CustomerName  string
	CustomerSurname  string
	Item      string
	Quantity  int
	Status    string
	Phone     string
	Address   string
	Email     string
	OrderTime string
	Price     float64
	Total     float64
	Reason    string

	MenuName         string
	MenuDescription  string
	MenuImage        string
} 