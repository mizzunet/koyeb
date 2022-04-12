package headless

import (
	"context"
	// "fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
)

func Scrape() string {
	chromeBin := os.Getenv("GOOGLE_CHROME_SHIM")

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
	var title string
	var loc string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://1lib.domains/?redirectUrl=/`),
		chromedp.Title(&title),
		chromedp.Location(&loc),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(title, loc)

	return strings.TrimSpace(res)
}