package main

import (
	"example.com/headless"
	"example.com/myip"
	"example.com/stock"
	"example.com/zlib"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

type Zlib struct {
	Query string `form:"q"`
}

type DATA struct {
	status, sensex, nifty string
}

func zlib_do(c *gin.Context) {
	ip := myip.GetIP()
	log.Println("Doing zlib, IP: ", ip)
	var f Zlib
	c.Bind(&f)
	o := zlib.DownloadBook(f.Query)
	c.JSON(200, gin.H{
		"name":  o.Name,
		"link":  o.Link,
		"error": o.Error,
	})
}
func headless_do(c *gin.Context) {
	log.Println("Doing headless")
	i := headless.Scrape()
	c.JSON(200, gin.H{
		"scrape": i,
	})
}
func stock_do(c *gin.Context) {
	log.Println("Doing stock")
	var i DATA
	i.status, i.sensex, i.nifty = stock.Parse()
	c.JSON(200, gin.H{
		// "hey":   "there",
		// "hello": i,
		"SENSEX": i.sensex,
		"NIFTY":  i.nifty,
		"STATUS": i.status,
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
	api.GET("/headless", headless_do)
	r.GET("/zlib", zlib_do)
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run()
}
