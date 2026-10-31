package repository

import (
	"food-ordering-system-project/config"
	"food-ordering-system-project/models"
)

func GetDashboardStats() (*models.DashboardStats, error) {

	var stats models.DashboardStats

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders",
	).Scan(&stats.TotalOrders)

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE status='PENDING'",
	).Scan(&stats.PendingOrders)

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE status='CLOSED'",
	).Scan(&stats.ClosedOrders)

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE status='CANCELLED'",
	).Scan(&stats.RejectedOrders)

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE DATE(orderTime)=CURDATE()",
	).Scan(&stats.TodayOrders)

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE status='PENDING' AND DATE(orderTime)=CURDATE()",
	).Scan(&stats.TodayPending)

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE status='CLOSED' AND DATE(orderTime)=CURDATE()",
	).Scan(&stats.TodayClosed)

	config.DB.QueryRow(
		"SELECT COUNT(*) FROM orders WHERE status='CANCELLED' AND DATE(orderTime)=CURDATE()",
	).Scan(&stats.TodayRejected)

	config.DB.QueryRow("SELECT IFNULL(SUM(total),0) FROM orders WHERE status='CLOSED'",).Scan(&stats.TotalSales)

	config.DB.QueryRow(`SELECT IFNULL(SUM(total),0) FROM orders WHERE status='CLOSED' AND DATE(orderTime)=CURDATE()`,
	).Scan(&stats.TodaySales)


	config.DB.QueryRow(`SELECT IFNULL(SUM(total),0)FROM orders WHERE status='CANCELLED'`,).Scan(&stats.TotalCancelledSales)

    config.DB.QueryRow(`SELECT IFNULL(SUM(total),0) FROM orders WHERE status='CANCELLED' AND DATE(orderTime)=CURDATE()`,).Scan(&stats.TodayCancelledSales)

	return &stats, nil
}