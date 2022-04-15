package main

import (
	// "example.com/headless"
	"example.com/myip"
	"example.com/stock"
	"example.com/zlibrary"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

type Zlib struct {
	Query string `form:"q"`
}

type DATA struct {
	date, status, sensex, nifty string
}

func zlib_do(c *gin.Context) {
	log.Println("Doing zlib, IP: ", myip.GetIP())
	var f Zlib
	c.Bind(&f)
	o := zlibrary.Query(f.Query)
	c.JSON(200, gin.H{
		"error":  o.Error,
		"url":    o.UploadURL,
		"name":   o.Name,
		"file":   o.FileName,
		"format": o.Format,
	})
}

// func headless_do(c *gin.Context) {
// u := headless.GetRedirectURL("https://1lib.domains/?redirectUrl=/")
// log.Println(u)
// log.Println("headless")
// i := headless.Scrape()
// c.JSON(200, gin.H{
// "scrape": i,
// })
// }
func stock_do(c *gin.Context) {
	log.Println("Doing stock")
	var i DATA
	i.date, i.status, i.sensex, i.nifty = stock.Parse()
	c.JSON(200, gin.H{
		// "hey":   "there",
		// "hello": i,
		"SENSEX": i.sensex,
		"NIFTY":  i.nifty,
		"STATUS": i.status,
		"DATE":   i.date,
	})
}
func main() {
	// fmt.Println(i.status)
	log.Println("Starting")
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	api := r.Group("/api")

	api.GET("/stock", stock_do)
	// api.GET("/headless", headless_do)
	r.GET("/zlib", zlib_do)
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run()
}
