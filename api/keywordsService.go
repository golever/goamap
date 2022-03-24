package api

import (
	"github.com/gin-gonic/gin"
	"goamap.mod/request"
	"net/http"
)

func KeyWords(c *gin.Context)  {
	k := c.Param("keywords")
	t := c.Param("types")
	p := request.Params{Keywords: k, Types: t}

	content := request.Send(p)
	c.String(http.StatusOK, content)
}