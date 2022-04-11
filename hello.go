package main

import (
	"fmt"
	"os"
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"    
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
    // "fmt"
)
type DATA struct {
	status, sensex, nifty  string
}

func main() {
    // fmt.Println(i.status)
    r := gin.Default()

    r.GET("/hello", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

    api := r.Group("/api")

    api.GET("/stock", func(c *gin.Context) {
	    i :=  getMarket()
        c.JSON(200, gin.H{
            "SENSEX": i.sensex,
            "NIFTY":  i.nifty,
            "STATUS": i.status,
        })
    })
    api.GET("/headless", func(c *gin.Context) {
	    i :=  headless()
        c.JSON(200, gin.H{
            "scrape": i,
        })
    })
    r.Use(static.Serve("/", static.LocalFile("./views", true)))

    r.Run()
}

func getMarket() DATA {
    var i DATA

	link := "https://economictimes.indiatimes.com/markets"
	c := colly.NewCollector()
	
	c.OnHTML("div.content_area",func(e *colly.HTMLElement) {
        // fmt.Println("Scraping")
		i.status =	e.DOM.Find("span.mktStatus").Text()
		i.sensex = e.DOM.Find("div[data-pos='1']").Find("span.change").Text()
		i.nifty	= e.DOM.Find("div[data-pos='2']").Find("span.change").Text()
	})
	
	c.Visit(link)
    return i
}

func headless() string {
	chromeBin := os.Getenv("GOOGLE_CHROME_SHIM")
	fmt.Println("chrome path: %+v", chromeBin)

	options := []chromedp.ExecAllocatorOption{
	        chromedp.ExecPath(chromeBin),
	        chromedp.Flag("headless", true),
	        chromedp.Flag("blink-settings", "imageEnable=false"),
	        chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko)`),
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

		// create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		chromedp.Text(`.Documentation-overview`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(res)
}
