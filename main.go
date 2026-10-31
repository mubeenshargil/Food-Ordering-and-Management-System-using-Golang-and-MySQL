package main

import (
	"food-ordering-system-project/config"
	"food-ordering-system-project/routes"
	"os"
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.Static("/static", "./static")
     
	r.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routes.SetupRoutes(r)

	port := os.Getenv("SERVER_PORT")

		if port == "" {
			log.Fatal("SERVER_PORT not configured")
		}

		err := r.Run(":" + port)

		if err != nil {
			log.Fatal(err)
		}
}