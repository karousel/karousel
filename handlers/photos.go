package handlers

import (
	"github.com/karousel/karousel/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetPhotoResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var photos []models.Photo
	db.Find(&photos)

	c.JSON(200, photos)
}
