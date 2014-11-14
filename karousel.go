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

var (
	config models.Configuration
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
	router.Use(middleware.Configure(config))

	router.POST("/tokens/", handlers.PostTokenResource)
	router.POST("/users/", handlers.PostUserResource)

	authenticated := router.Group("/")

	authenticated.Use(middleware.Authenticate())
	{
		authenticated.GET("/users/", handlers.GetUserResource)
		authenticated.GET("/users/:id/", handlers.GetUserInstance)
		authenticated.DELETE("/users/:id/", handlers.DeleteUserInstance)

		authenticated.GET("/collections/", handlers.GetCollectionResource)
		authenticated.POST("/collections/", handlers.PostCollectionResource)
		authenticated.GET("/collections/:id/", handlers.GetCollectionInstance)
		authenticated.DELETE("/collections/:id/", handlers.DeleteCollectionInstance)

		authenticated.GET("/albums/", handlers.GetAlbumResource)
		authenticated.POST("/albums/", handlers.PostAlbumResource)
		authenticated.GET("/albums/:id/", handlers.GetAlbumInstance)
		authenticated.DELETE("/albums/:id/", handlers.DeleteAlbumInstance)
	}

	router.Run(fmt.Sprintf(":%v", config.Web.Port))
}
