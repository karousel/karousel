package handlers

import (
	"github.com/citruspi/karousel/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAlbumResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var albums []models.Album
	db.Find(&albums)

	c.JSON(200, albums)
}
