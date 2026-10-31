package repository

import (
	"database/sql"
	"food-ordering-system-project/models"
)

type UserRepository struct {
	DB *sql.DB
}

type User struct {
	ID       int
	Role     string
	Name     string
	Surname  string
	Password string
}

// func GetUserByEmail(email string) (*User, error) {
func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User

	err := r.DB.QueryRow(
		`SELECT
		id,
		role,
		name,
		surname,
		password
		FROM user
		WHERE email=? AND status=1`,
		email,
	).Scan(
		&user.ID,
		&user.Role,
		&user.Name,
		&user.Surname,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// func CreateUser(
func (r *UserRepository) CreateUser(
	name string,
	surname string,
	phone string,
	email string,
	password string,
	plainPassword string,
	role string,
	status string,
) error {

	_, err := r.DB.Exec(
		`INSERT INTO user
		(name,surname,phone,email,password,plainpassword,role,status)
		VALUES(?,?,?,?,?,?,?,?)`,
		name,
		surname,
		phone,
		email,
		password,
		plainPassword,
		role,
		status,
	)

	return err
}


// func GetAllUsers() ([]models.User, error) {
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query(
		`SELECT
		id,
		name,
		surname,
		phone,
		email,
		role,
		status
		FROM user
		ORDER BY id DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Phone,
			&user.Email,
			&user.Role,
			&user.Status,
		)

		users = append(users, user)
	}

	return users, nil
}

// func UpdateUser(
func (r *UserRepository) UpdateUser(
	id string,
	name string,
	surname string,
	phone string,
	email string,
	role string,
	status string,
) error {

	_, err := r.DB.Exec(
		`UPDATE user
		SET
		name=?,
		surname=?,
		phone=?,
		email=?,
		role=?,
		status=?
		WHERE id=?`,
		name,
		surname,
		phone,
		email,
		role,
		status,
		id,
	)

	return err
}