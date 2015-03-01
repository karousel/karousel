package handlers

import (
	"fmt"
	"time"

	"github.com/karousel/karousel/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAlbumResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var albums []models.Album
	db.Find(&albums)

	c.JSON(200, albums)
}

func PostAlbumResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var album models.Album

	c.Bind(&album)

	if (album.Name == "") || (album.CollectionId == 0) {
		response := make(map[string]string)
		response["error"] = "Incomplete submission."
		c.JSON(400, response)
	} else {
		var queryCollection models.Collection

		db.First(&queryCollection, album.CollectionId)

		if queryCollection.Name == "" {
			response := make(map[string]string)
			response["error"] = "Resource not found."
			c.JSON(404, response)
		} else {
			var queryAlbums []models.Album
			var duplicate bool

			db.Where("name = ?", album.Name).Find(&queryAlbums)

			for _, queryAlbum := range queryAlbums {
				if queryAlbum.CollectionId == album.CollectionId {
					duplicate = true
					break
				}
			}

			if duplicate {
				response := make(map[string]string)
				response["error"] = "Duplicate resource."
				c.JSON(409, response)
			} else {
				album.Created = time.Now().UTC()
				db.Create(&album)

				locationHeader := fmt.Sprintf("/albums/%v", album.Id)

				c.Writer.Header().Set("Location", locationHeader)
			}
		}
	}
}

func GetAlbumInstance(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	id := c.Params.ByName("id")

	var album models.Album

	db.First(&album, id)

	if album.Name == "" {
		response := make(map[string]string)
		response["error"] = "Resource not found."
		c.JSON(404, response)
	} else {
		c.JSON(200, album)
	}
}

func DeleteAlbumInstance(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)
	consumer := c.MustGet("consumer").(models.User)

	id := c.Params.ByName("id")

	var album models.Album

	db.First(&album, id)

	if album.Name == "" {
		response := make(map[string]string)
		response["error"] = "Resource not found."
		c.JSON(404, response)
	} else {
		if consumer.Admin {
			db.Delete(&album)
			c.JSON(200, album)
		} else {
			response := make(map[string]string)
			response["error"] = "Invalid credentials."
			c.JSON(401, response)
		}
	}
}
