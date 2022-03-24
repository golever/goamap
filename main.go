package main

import "goamap.mod/db"

func main() {

	//r := gin.Default()
	//r.GET("/:keywords/:types", func(c *gin.Context) {
	//	api.KeyWords(c)
	//})
	//r.Run(":8080")
	db.Insert()
}
