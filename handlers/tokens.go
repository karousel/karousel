package handlers

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/karousel/karousel/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func PostTokenResource(c *gin.Context) {
	db := c.MustGet("db").(gorm.DB)
	config := c.MustGet("config").(models.Configuration)

	var user models.User

	c.Bind(&user)

	if (user.Username == "") || (user.Password == "") {
		response := make(map[string]string)
		response["error"] = "Incomplete submission."
		c.JSON(400, response)
	} else {
		var queryUser models.User

		db.Where("username = ?", user.Username).First(&queryUser)

		if queryUser.Username == "" {
			response := make(map[string]string)
			response["error"] = "Resource not found."
			c.JSON(404, response)
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(queryUser.Password), []byte(user.Password))

			if err != nil {
				response := make(map[string]string)
				response["error"] = "Invalid credentials."
				c.JSON(401, response)
			} else {
				key, err := ioutil.ReadFile(config.Keys.Private)

				if err != nil {
					log.Fatal(err)
				}

				token := jwt.New(jwt.GetSigningMethod("RS256"))

				token.Claims["id"] = queryUser.Id

				expires := time.Now().Add(time.Hour * 72).UTC()

				token.Claims["life"] = expires.Format(time.RFC3339)

				tokenString, err := token.SignedString(key)

				response := make(map[string]string)
				response["token"] = tokenString
				response["expires"] = expires.String()

				c.JSON(200, response)
			}
		}
	}
}
