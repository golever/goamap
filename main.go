package main

import (
	"github.com/gin-gonic/gin"
	"goamap.mod/services"
)

func main() {

	r := gin.Default()
	r.GET("/:keywords/:types", func(c *gin.Context) {
		services.KeyWords(c)
	})
	r.Run(":8080")

}
