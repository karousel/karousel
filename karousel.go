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

	router.GET("/users/", middleware.Authenticate(), handlers.GetUserResource)
	router.POST("/users/", handlers.PostUserResource)
	router.GET("/users/:id/", middleware.Authenticate(), handlers.GetUserInstance)
	router.DELETE("/users/:id/", middleware.Authenticate(), handlers.DeleteUserInstance)

	router.GET("/collections/", middleware.Authenticate(), handlers.GetCollectionResource)
	router.POST("/collections/", middleware.Authenticate(), handlers.PostCollectionResource)
	router.GET("/collections/:id/", middleware.Authenticate(), handlers.GetCollectionInstance)
	router.DELETE("/collections/:id/", middleware.Authenticate(), handlers.DeleteCollectionInstance)

	router.GET("/albums/", middleware.Authenticate(), handlers.GetAlbumResource)
	router.POST("/albums/", middleware.Authenticate(), handlers.PostAlbumResource)
	router.GET("/albums/:id/", middleware.Authenticate(), handlers.GetAlbumInstance)
	router.DELETE("/albums/:id/", middleware.Authenticate(), handlers.DeleteAlbumInstance)

	router.Run(fmt.Sprintf(":%v", config.Web.Port))
}
