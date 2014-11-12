package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/citruspi/Karousel-API/handlers"
	"github.com/citruspi/Karousel-API/middleware"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Web struct {
		Port string
	}
	Database struct {
		Host string
		Port string
		Name string
	}
}

var (
	config  Configuration
	session *r.Session
	err     error
)

func main() {
	data, err := ioutil.ReadFile("karousel.conf")

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		log.Fatal(err)
	}

	session, err = r.Connect(r.ConnectOpts{
		Address:  fmt.Sprintf("%v:%v", config.Database.Host, config.Database.Port),
		Database: config.Database.Name,
	})

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Database(session))

	router.GET("/users/", handlers.GetUserResource)

	router.Run(fmt.Sprintf(":%v", config.Web.Port))
}
