package handlers

import (
	"fmt"
	"time"

	"github.com/citruspi/karousel/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PostUserResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var user models.User

	c.Bind(&user)

	if (user.Username == "") || (user.Email == "") || (user.Password == "") {
		response := make(map[string]string)
		response["error"] = "Incomplete submission."
		c.JSON(400, response)
	} else {
		var queryUser models.User

		db.Where("username = ?", user.Username).First(&queryUser)

		if queryUser.Username != "" {
			response := make(map[string]string)
			response["error"] = "Duplicate resource."
			c.JSON(409, response)
		} else {
			db.Where("email = ?", user.Email).First(&queryUser)

			if queryUser.Username != "" {
				response := make(map[string]string)
				response["error"] = "Duplicate resource."
				c.JSON(409, response)
			} else {
				user.Joined = time.Now().UTC()

				db.Create(&user)

				user.Password = ""

				locationHeader := fmt.Sprintf("/users/%v", user.Id)

				c.Writer.Header().Set("Location", locationHeader)
			}
		}
	}
}

func GetUserResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)

	var users []models.User
	db.Find(&users)

	c.JSON(200, users)
}
