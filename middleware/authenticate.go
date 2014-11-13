package middleware

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/citruspi/karousel/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(gorm.DB)

		tokenString := c.Request.Header.Get("X-Authentication-Token")

		if tokenString == "" {
			response := make(map[string]string)
			response["error"] = "Invalid credentials."
			c.JSON(401, response)
			c.Abort(401)
		} else {
			token, err := jwt.Parse(tokenString, verify)

			if err != nil {
				log.Fatal(err)
			} else {
				if token.Valid {

					now := time.Now().UTC()
					expiration, err := time.Parse(time.RFC3339, token.Claims["life"].(string))

					if err != nil {
						log.Fatal(err)
					}

					if now.After(expiration) {
						response := make(map[string]string)
						response["error"] = "Invalid credentials."
						c.JSON(401, response)
						c.Abort(401)
					}

					var user models.User

					db.First(&user, token.Claims["id"])

					if user.Username == "" {
						response := make(map[string]string)
						response["error"] = "Someone fucked up."
						c.JSON(500, response)
						c.Abort(500)
					} else {
						c.Set("consumer", user)
						c.Next()
					}
				} else {
					response := make(map[string]string)
					response["error"] = "Someone fucked up."
					c.JSON(500, response)
					c.Abort(500)
				}
			}
		}
	}
}

func verify(*jwt.Token) (interface{}, error) {
	key, err := ioutil.ReadFile("key.pub")

	if err != nil {
		log.Fatal(err)
	}

	return key, nil
}
