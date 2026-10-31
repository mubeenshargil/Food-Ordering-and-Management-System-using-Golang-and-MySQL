package repository

import (
	"database/sql"
	"food-ordering-system-project/models"
)

type MenuRepository struct {
	DB *sql.DB
}

// func GetMenus() ([]models.Menu, error) {
func (r *MenuRepository) GetMenus() ([]models.Menu, error) {

	// rows, err := config.DB.Query(
	rows, err := r.DB.Query(

		"SELECT id,name,description,price,image,status FROM menu",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var menus []models.Menu

	for rows.Next() {

		var m models.Menu

		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.Description,
			&m.Price,
			&m.Image,
			&m.Status,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, m)
	}

	return menus, nil
}

// func GetMenuByID(id string) (*models.Menu, error) {
func (r *MenuRepository) GetMenuByID(id string) (*models.Menu, error) {

	var menu models.Menu

	err := r.DB.QueryRow(
		"SELECT id,name,description,price,image FROM menu WHERE id=?",
		id,
	).Scan(
		&menu.ID,
		&menu.Name,
		&menu.Description,
		&menu.Price,
		&menu.Image,
	)

	if err != nil {
		return nil, err
	}

	return &menu, nil
}

// func CreateMenu(
func (r *MenuRepository) CreateMenu(
	name string,
	description string,
	price string,
	image string,
	status string,
) error {

	_, err := r.DB.Exec(
		"INSERT INTO menu(name,description,price,image,status) VALUES(?,?,?,?,?)",
		name,
		description,
		price,
		image,
		status,
	)

	return err
}
// func UpdateMenu(
func (r *MenuRepository) UpdateMenu(
	id string,
	name string,
	description string,
	price string,
	status string,
) error {

	_, err := r.DB.Exec(
		`UPDATE menu
		SET Name=?,
		Description=?,
		Price=?,
		Status=?
		WHERE ID=?`,
		name,
		description,
		price,
		status,
		id,
	)

	return err
}