package handlers

import (
	"fmt"
	"time"

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

func PostCollectionResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var collection models.Collection

	c.Bind(&collection)

	if collection.Name == "" {
		response := make(map[string]string)
		response["error"] = "Incomplete submission."
		c.JSON(400, response)
	} else {
		var queryCollection models.Collection

		db.Where("name = ?", collection.Name).First(&queryCollection)

		if queryCollection.Name != "" {
			response := make(map[string]string)
			response["error"] = "Duplicate resource."
			c.JSON(409, response)
		} else {
			collection.Created = time.Now().UTC()

			db.Create(&collection)

			locationHeader := fmt.Sprintf("/collections/%v", collection.Id)

			c.Writer.Header().Set("Location", locationHeader)
		}
	}
}

func GetCollectionInstance(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	id := c.Params.ByName("id")

	var collection models.Collection

	db.First(&collection, id)

	if collection.Name == "" {
		response := make(map[string]string)
		response["error"] = "Resource not found."
		c.JSON(404, response)
	} else {
		c.JSON(200, collection)
	}
}
