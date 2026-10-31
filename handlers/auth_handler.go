package handlers

import (
	"food-ordering-system-project/services"
	"net/http"
     "golang.org/x/crypto/bcrypt"
	 
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserService *services.UserService
}

// func Register(c *gin.Context) {
func (h *AuthHandler) Register(c *gin.Context) {
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	plainpassword := c.PostForm("password")
	status := c.PostForm("status")
	role := c.PostForm("role")


	password := c.PostForm("password")

	hashedPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)

	err := h.UserService.CreateUser(
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

	c.Redirect(http.StatusSeeOther, "login")
}

// func Login(c *gin.Context) {
func (h *AuthHandler) Login(c *gin.Context) {

	email := c.PostForm("email")

	password := c.PostForm("password")

 	user, err := h.UserService.GetUserByEmail(email)

	if err != nil {

		c.String(401, "Invalid Email")
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {

		c.String(401, "Invalid Password")
		return
	}

	session := sessions.Default(c)

	session.Set("user_id", user.ID)
	session.Set("role", user.Role)
	session.Set("name", user.Name)
	session.Set("surname", user.Surname)


	session.Save()

		if user.Role == "admin" {

		c.Redirect(http.StatusSeeOther,"/admin",)

		} else if user.Role == "cashier" {

			c.Redirect(http.StatusSeeOther,"/cashier/orders",)

		} else {

			c.Redirect(http.StatusSeeOther,"/",)
		}
	
}

// func Logout(c *gin.Context) {
func (h *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusSeeOther, "login")
}