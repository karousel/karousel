package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/citruspi/karousel/handlers"
	"github.com/citruspi/karousel/middleware"
	"github.com/citruspi/karousel/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Web struct {
		Port string
	}
}

var (
	config Configuration
	err    error
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

	db, err := gorm.Open("sqlite3", "karousel.db")

	db.DB()
	db.DB().Ping()

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Album{}, &models.Collection{})

	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Database(db))

	router.GET("/users/", handlers.GetUserResource)

	router.Run(fmt.Sprintf(":%v", config.Web.Port))
}
