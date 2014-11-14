package handlers

import (
	"github.com/citruspi/karousel/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetCollectionResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var collections []models.Collection
	db.Find(&collections)

	c.JSON(200, collections)
}
