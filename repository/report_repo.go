package repository

import (
	"database/sql"
	"food-ordering-system-project/models"
)

type ReportRepository struct {
	DB *sql.DB
}

// func GetSalesReport(fromDate string,toDate string,) (models.SalesReport, error) {
func (r *ReportRepository) GetSalesReport(fromDate string,toDate string,) (models.SalesReport, error) {
	var report models.SalesReport

	err := r.DB.QueryRow(
		`SELECT
		COUNT(*) AS totalOrders,
		IFNULL(SUM(total),0) AS totalRevenue,
		IFNULL(SUM(quantity),0) AS totalQuantity,
		IFNULL(AVG(total),0) AS averageOrderValue

		FROM orders

		WHERE DATE(orderTime)
		BETWEEN ? AND ?`,
		fromDate,
		toDate,
	).Scan(
		&report.TotalOrders,
		&report.TotalRevenue,
		&report.TotalQuantity,
		&report.AverageOrderValue,
	)

	return report, err
}

// func GetMenuPerformanceReport() ([]models.MenuPerformance, error) {
func (r *ReportRepository) GetMenuPerformanceReport() ([]models.MenuPerformance, error){

	rows, err := r.DB.Query(
		`SELECT
		m.name,
		COUNT(o.id) AS totalOrders,
		IFNULL(SUM(o.quantity),0) AS totalQuantity,
		IFNULL(SUM(o.total),0) AS revenue

		FROM menu m

		LEFT JOIN orders o
		ON m.id = o.MenuID

		GROUP BY m.id, m.name

		ORDER BY revenue DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var report []models.MenuPerformance

	for rows.Next() {

		var item models.MenuPerformance

		rows.Scan(
			&item.MenuName,
			&item.TotalOrders,
			&item.TotalQuantity,
			&item.Revenue,
		)

		report = append(report, item)
	}

	return report, nil
}

// func GetCustomerOrdersReport() ([]models.CustomerOrderReport, error) {
func (r *ReportRepository) GetCustomerOrdersReport() ([]models.CustomerOrderReport, error){

	rows, err := r.DB.Query(
		`SELECT
		CONCAT(
			customer_name,
			' ',
			customer_surname
		) AS customerName,

		email,

		COUNT(*) AS totalOrders,

		IFNULL(
			SUM(total),
			0
		) AS totalSpending

		FROM orders

		GROUP BY
		customer_name,
		customer_surname,
		email

		ORDER BY totalSpending DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var report []models.CustomerOrderReport

	for rows.Next() {

		var item models.CustomerOrderReport

		rows.Scan(
			&item.CustomerName,
			&item.Email,
			&item.TotalOrders,
			&item.TotalSpending,
		)

		report = append(report, item)
	}

	return report, nil
}

// func GetOrdersReport(startDate string,endDate string,status string) ([]models.OrdersReport, error) {
func (r *ReportRepository) GetOrdersReport(startDate string,endDate string,status string) ([]models.OrdersReport, error) {
	query := `
	SELECT
	id,

	CONCAT(
		customer_name,
		' ',
		customer_surname
	),

	email,

	status,

	IFNULL(reason,''),

	total,

	DATE_FORMAT(
		orderTime,
		'%d %b %Y %h:%i %p'
	)

	FROM orders

	WHERE 1=1
	`

	var args []interface{}

	if startDate != "" && endDate != "" {

		query += `
		AND DATE(orderTime)
		BETWEEN ? AND ?
		`

		args = append(
			args,
			startDate,
			endDate,
		)
	}

	if status != "" {

		query += `
		AND status=?
		`

		args = append(
			args,
			status,
		)
	}

	query += `
	ORDER BY id DESC
	`

	rows, err := r.DB.Query(
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var report []models.OrdersReport

	for rows.Next() {

		var item models.OrdersReport

		err := rows.Scan(
			&item.OrderID,
			&item.CustomerName,
			&item.Email,
			&item.Status,
			&item.Reason,
			&item.Total,
			&item.OrderTime,
		)

		if err != nil {
			return nil, err
		}

		report = append(report, item)
	}

	return report, nil
}