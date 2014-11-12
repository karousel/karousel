package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/citruspi/Karousel-API/models"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
)

func GetUserResource(c *gin.Context) {
	session := c.MustGet("session").(*r.Session)

	rows, err := r.Table("users").Run(session)

	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	err = rows.All(&users)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, users)
}
