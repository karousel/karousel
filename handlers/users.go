package handlers

import (
	"github.com/citruspi/karousel/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetUserResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var users []models.User
	db.Find(&users)

	c.JSON(200, users)
}
