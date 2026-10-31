package handlers

import (
	"net/http"
	"food-ordering-system-project/helpers"
	"food-ordering-system-project/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserService *services.UserService
}

// func AdminUsersPage(c *gin.Context) {
func (h *UserHandler) AdminUsersPage(c *gin.Context) {
	
	users, err := h.UserService.GetAllUsers()

	if err != nil {

		c.String(500, err.Error())
		return
	}

	user := helpers.GetSessionData(c)


	c.HTML(http.StatusOK, "ad_users.html", gin.H{
		"title": "Users",
		"users":    users,
		"role": user.Role,
		"username": user.Username,
		"name":       user.Name,
		"Surname":       user.Surname,
	})
}

// func AddUser(c *gin.Context) {
func (h *UserHandler) AddUser(c *gin.Context) {

	name := c.PostForm("name")

	surname := c.PostForm("surname")

	phone := c.PostForm("phone")

	email := c.PostForm("email")

	plainpassword := c.PostForm("password")

	password := c.PostForm("password")

	status := c.PostForm("status")

	

	hashedPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)

	role := c.PostForm("role")

		err :=  h.UserService.CreateUser(
		name,
		surname,
		phone,
		email,
		string(hashedPassword),
		plainpassword,
		role,
		status,
	)

	if err != nil {

		c.String(500, err.Error())
		return
	}

	c.Redirect(302, "/admin/users")
}

// func UpdateUser(c *gin.Context) {
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	name := c.PostForm("name")

	surname := c.PostForm("surname")

	phone := c.PostForm("phone")

	email := c.PostForm("email")

	role := c.PostForm("role")

	status := c.PostForm("status")

	err := h.UserService.UpdateUser(
	id,
	name,
	surname,
	phone,
	email,
	role,
	status,
)

	if err != nil {

		c.String(500, err.Error())
		return
	}

	c.Redirect(302, "/admin/users")
}

// func DeleteUser(c *gin.Context) {

// 	id := c.Param("id")

// 	_, err := config.DB.Exec(
// 		"DELETE FROM users WHERE id=?",
// 		id,
// 	)

// 	if err != nil {

// 		c.String(500, err.Error())
// 		return
// 	}

// 	c.Redirect(302, "/admin/users")
// }