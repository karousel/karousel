package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/citruspi/karousel/handlers"
	"github.com/citruspi/karousel/middleware"
	"github.com/citruspi/karousel/models"

	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
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

	db, err := sql.Open("sqlite3", "karousel.db")

	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(models.Album{}, "albums").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.Collection{}, "collections").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.Photo{}, "photos").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()

	if err != nil {
		log.Fatal(err)
	}

	defer dbmap.Db.Close()

	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Database(dbmap))

	router.GET("/users/", handlers.GetUserResource)

	router.Run(fmt.Sprintf(":%v", config.Web.Port))
}
