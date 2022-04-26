package main

import (
	// "example.com/headless"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/stock"
	"example.com/zlibrary"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Zlib struct {
	Query string `form:"q"`
}

type Stock struct {
	Date, Status, Sensex, Nifty string
}

var ip string

func getIP() string {
	url := "https://api.ipify.org?format=text" // we are using a pulib IP API, we're using ipify here, below are some others

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(ip)
}
func zlib_do(c *gin.Context) {
	var f Zlib
	c.Bind(&f)
	o := zlibrary.Query(f.Query)
	c.JSON(200, gin.H{
		"IP":    getIP(),
		"ERROR": o.Error,
		"URL":   o.UploadURL,
		"NAME":  o.Name,
		"FILE":  o.FileName,
		"SIZE":  o.Size,
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
func getStock(c *gin.Context) {
	log.Println("Doing stock")
	var i Stock
	i.Date, i.Status, i.Sensex, i.Nifty = stock.Parse()
	c.JSON(200, gin.H{
		"IP":     getIP(),
		"SENSEX": i.Sensex,
		"NIFTY":  i.Nifty,
		"STATUS": i.Status,
		"DATE":   i.Date,
	})
}
func main() {
	log.Println("Starting")
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// api := r.Group("/api")

	// api.GET("/stock", getStock)
	r.GET("/stock", getStock)
	// api.GET("/headless", headless_do)
	r.GET("/zlib", zlib_do)

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run()
}
