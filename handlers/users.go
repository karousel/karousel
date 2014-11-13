package handlers

import (
	"log"

	"github.com/citruspi/karousel/models"

	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
)

func GetUserResource(c *gin.Context) {
	db := c.MustGet("db").(*gorp.DbMap)

	var users []models.User
	_, err := db.Select(&users, "select * from users order by id")

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, users)
}
