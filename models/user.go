package models

type User struct {
	ID       int
	Name     string
	Surname  string
	Phone    string
	Email    string
	Password string
	PlainPassword string
	Role     string
	Status    int
}