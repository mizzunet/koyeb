package myip

import (
	"github.com/gocolly/colly/v2"
	// "io/ioutil"
	"log"
	// "net/http"
)

func GetIP() string {
	var ip string
	c := colly.NewCollector()

	c.OnHTML("input[name='ip']", func(e *colly.HTMLElement) {
		ip, _ := e.DOM.Attr("value")
		log.Println(ip)
	})
	c.Visit("https://ip.me")
	return ip
}