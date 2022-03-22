package main

import (
	"github.com/gin-gonic/gin"
	"goamap.mod/request"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/:keywords/:types", func(c *gin.Context) {
		k := c.Param("keywords")
		t := c.Param("types")
		p := request.Params{Keywords: k, Types: t}

		content := request.Send(p)
		c.String(http.StatusOK, content)
	})
	r.Run(":8080")
}
