package main

import (
	"log"

	"github.com/citruspi/Karousel-API/handlers"
	"github.com/citruspi/Karousel-API/middleware"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
)

var (
	session *r.Session
	err     error
)

func main() {
	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "karousel",
	})

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Database(session))

	router.GET("/users/", handlers.GetUserResource)

	router.Run(":8000")
}
