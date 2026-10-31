package models

type DashboardStats struct {
	TotalOrders    int
	PendingOrders  int
	ClosedOrders   int
	RejectedOrders int

	TodayOrders    int
	TodayPending   int
	TodayClosed    int
	TodayRejected  int

	TotalSales     float64
	TodaySales     float64

	TotalCancelledSales float64
	TodayCancelledSales float64
}