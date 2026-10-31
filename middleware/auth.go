package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)

		role := session.Get("role")

		if role != "admin" {

			// c.String(http.StatusUnauthorized, "Admin Access Only")
			// c.Abort()
			// return
			c.Redirect(http.StatusFound, "/login")
				c.Abort()
				return
		}

		c.Next()
	}
}

func CashierRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)

		role := session.Get("role")

		if role != "cashier" {

			c.Redirect(
				http.StatusSeeOther,
				"/login",
			)

			c.Abort()
			return
		}

		c.Next()
	}
}