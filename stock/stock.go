package stock

import (
	"github.com/gocolly/colly/v2"
)

type DATA struct {
	status, sensex, nifty string
}

func Parse() (string, string, string) {
	var i DATA

	link := "https://economictimes.indiatimes.com/markets"
	c := colly.NewCollector()

	c.OnHTML("div.content_area", func(e *colly.HTMLElement) {
		// fmt.Println("Scraping")
		i.status = e.DOM.Find("span.mktStatus").Text()
		i.sensex = e.DOM.Find("div[data-pos='1']").Find("span.change").Text()
		i.nifty = e.DOM.Find("div[data-pos='2']").Find("span.change").Text()
	})

	c.Visit(link)
	return i.status, i.sensex, i.nifty
}