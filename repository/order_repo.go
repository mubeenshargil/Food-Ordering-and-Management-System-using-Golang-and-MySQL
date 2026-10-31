package repository

import (
	"database/sql"
	"food-ordering-system-project/models"
	
)


type OrderRepository struct {
	DB *sql.DB
}


// func CreateOrder(
func (r *OrderRepository) CreateOrder(
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

	// result, err := config.DB.Exec(
	result, err := r.DB.Exec(
		`INSERT INTO orders
		(customer_name,customer_surname,phone,address,email,MenuID,Quantity,Status,price,total)
		VALUES(?,?,?,?,?,?,?,?,?,?)`,
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

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// func CreateOrderHistory(
func (r *OrderRepository) CreateOrderHistory(
	orderID int64,
	status string,
) error {

	// _, err := config.DB.Exec(
	_, err := r.DB.Exec(
		`INSERT INTO orderhistory
		(orderId,status)
		VALUES(?,?)`,
		orderID,
		status,
	)

	return err
}

// func GetAllOrders() ([]models.Order, error) {
func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {

	// rows, err := config.DB.Query(`
	rows, err := r.DB.Query(`
	SELECT orders.id,
	CONCAT(
		orders.customer_name,
		' ',
		orders.customer_surname
	) AS customer_full_name,
	menu.name,
	orders.quantity,
	orders.status,
	orders.phone,
	orders.address,
	orders.email,
	DATE_FORMAT(orders.orderTime,'%d %b %Y %h:%i %p') as orderTime,
	orders.price,
	orders.total
	FROM orders
	JOIN menu
	ON orders.MenuID = menu.id
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []models.Order

	for rows.Next() {

		var order models.Order

		rows.Scan(
			&order.ID,
			&order.CustomerName,
			&order.Item,
			&order.Quantity,
			&order.Status,
			&order.Phone,
			&order.Address,
			&order.Email,
			&order.OrderTime,
			&order.Price,
			&order.Total,
		)

		orders = append(orders, order)
	}

	return orders, nil
}

// func UpdateOrderStatus(
func (r *OrderRepository) UpdateOrderStatus(
	id string,
	status string,
	updatedBy int,
	reason string,
) error {

	// _, err := config.DB.Exec(
	_, err := r.DB.Exec(

		"UPDATE orders SET status=?,reason=? WHERE id=?",
		status,
		reason,
		id,
	)

		// _, err = config.DB.Exec(
		_, err = r.DB.Exec(

		`INSERT INTO orderhistory
		(orderId,status,updatedBy,reason)
		VALUES(?,?,?,?)`,
		id,
		status,
		updatedBy,
		reason,
	)

	return err
}

// func GetOrderByID(id string) (models.Order, error) {
func (r *OrderRepository) GetOrderByID(id string) (models.Order, error) {

	var order models.Order

	// err := config.DB.QueryRow(
	err := r.DB.QueryRow(
	`SELECT
	orders.id,
	orders.customer_name,
	orders.customer_surname,
	orders.phone,
	orders.address,
	orders.email,
	orders.status,
	orders.price,
	orders.total,
	IFNULL(orders.reason,''),
    orders.quantity,
    DATE_FORMAT(orders.orderTime,'%d %b %Y %h:%i %p'),
	menu.name,
	menu.description,
	menu.image

	FROM orders

	JOIN menu
	ON orders.MenuID = menu.id

	WHERE orders.id=?`,
	id,
).Scan(
	&order.ID,
	&order.CustomerName,
	&order.CustomerSurname,
	&order.Phone,
	&order.Address,
	&order.Email,
	&order.Status,
	&order.Price,
	&order.Total,
	&order.Reason,
	&order.Quantity,
    &order.OrderTime,

	&order.MenuName,
	&order.MenuDescription,
	&order.MenuImage,
)

	return order, err
}

// func GetOrderHistory(orderID string) ([]models.OrderHistory, error) {
func (r *OrderRepository) GetOrderHistory(orderID string) ([]models.OrderHistory, error) {

	// rows, err := config.DB.Query(
	rows, err := r.DB.Query(
		`SELECT
			oh.id,
			oh.status,
			IFNULL(oh.reason,'') AS reason,
			IFNULL(
				CONCAT(u.name,' ',u.surname),
				'Customer'
			) AS updatedBy,
			DATE_FORMAT(
				oh.updatedAt,
				'%d %b %Y %h:%i %p'
			) AS updatedAt

			FROM orderhistory oh

			LEFT JOIN user u
			ON oh.updatedBy = u.id

			WHERE oh.orderId = ?

			ORDER BY oh.id ASC`,
		orderID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var history []models.OrderHistory

	for rows.Next() {

		var h models.OrderHistory

		err := rows.Scan(
			&h.ID,
			&h.Status,
			&h.Reason,
			&h.UpdatedBy,
			&h.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		history = append(history, h)
	}

	return history, nil
}