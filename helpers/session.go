package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	UserID   int
	Role     interface{}
	Name     interface{}
	Surname  interface{}
	Username interface{}
}

func GetSessionData(c *gin.Context) SessionData {

	session := sessions.Default(c)

	return SessionData{
		UserID:   session.Get("user_id").(int),
		Role:     session.Get("role"),
		Name:     session.Get("name"),
		Surname:  session.Get("surname"),
		Username: session.Get("username"),
	}
}