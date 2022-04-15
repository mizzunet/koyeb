package stock

import (
	// "fmt"
	"github.com/gocolly/colly/v2"
)

type DATA struct {
	date, status, sensex, nifty string
}

func Parse() (string, string, string, string) {
	var i DATA

	link := "https://economictimes.indiatimes.com/markets"
	c := colly.NewCollector()

	c.OnHTML("div.content_area", func(e *colly.HTMLElement) {
		// fmt.Println("Scraping")
		i.date = e.DOM.Find("span.date_format:nth-child(4)").Text()
		i.status = e.DOM.Find("span.mktStatus").Text()
		i.sensex = e.DOM.Find("div[data-pos='1']").Find("span.change").Text()
		i.nifty = e.DOM.Find("div[data-pos='2']").Find("span.change").Text()
	})

	c.Visit(link)
	return i.date, i.status, i.sensex, i.nifty
}
