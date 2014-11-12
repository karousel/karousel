package main

import (
	"github.com/citruspi/Karousel-API/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.Run(":8000")
}
