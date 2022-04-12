package main

import (
	"example.com/headless"
	"example.com/stock"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	// "fmt"
)

type DATA struct {
	status, sensex, nifty string
}

func main() {
	// fmt.Println(i.status)
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	api := r.Group("/api")

	api.GET("/stock", func(c *gin.Context) {
		var i DATA
		i.status, i.sensex, i.nifty = stock.Parse()
		c.JSON(200, gin.H{
			// "hey":   "there",
			// "hello": i,
			"SENSEX": i.sensex,
			"NIFTY":  i.nifty,
			"STATUS": i.status,
		})
	})
	api.GET("/headless", func(c *gin.Context) {
		i := headless.Scrape()
		c.JSON(200, gin.H{
			"scrape": i,
		})
	})
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run()
}
