package models

type OrderHistory struct {
	ID        int
	OrderId   int
	Status    string
	Reason    string
	UpdatedBy string
	UpdatedAt string

}