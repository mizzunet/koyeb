package zlib

import (
	// "fmt"
	"example.com/headless"

	// "github.com/root-gg/plik/plik"
	"log"
	"os"
	// "strings"

	"example.com/myip"
	"example.com/upload"
	"gopkg.in/headzoo/surf.v1"
)

type Output struct {
	Name  string
	URL   string
	Error string
	IP    string
}

var ret Output

func prepareFile(author, name string) (*os.File, string) {

	path := "./download/"
	fullpath := path + author + " - " + name + ".epub"

	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println(path, " not found, making it.")
		os.Mkdir(path, 0755)
	}
	if _, err := os.Stat(fullpath); !os.IsNotExist(err) {
		log.Println("File already been downloaded at ", fullpath)
	}

	file, err := os.Create(fullpath)
	if err != nil {
		log.Println("Failed to create file \n", err)
		ret.Error = "Failed to create file"
		// return ret

		log.Println("Folder: ", path)

	}
	return file, fullpath
}

func DownloadBook(query string) Output {
	filters := "?extensions[]=epub"
	// query := strings.TrimSpace(strings.Join(q, " "))
	if query == "" {
		log.Println("Search for something!")
		ret.Error = "Search for something!"
		return ret
	}

	base := headless.GetRedirectURL("https://1lib.domains/?redirectUrl=/")
	if base == "https://1lib.domains/?redirectUrl=/" {
		base = "https://u1lib.org/"
	}
	queryURL := base + "s/" + query + filters
	log.Println("Querying ", query)
	bow := surf.NewBrowser()
	url := queryURL

	// open search url
	log.Println("Visiting ", queryURL)
	err := bow.Open(url)
	if err != nil {
		log.Println("Failed to open link", url, err)
		ret.Error = "Failed to open link" + url
		return ret
	}
	// exit if no resule
	if bow.Find("div.notFound").Text() == "On your request nothing has been found. Do you want to create a ZAlert on this query?" {
		log.Println("No books found for ", query)
		ret.Error = "No books found for"
		return ret
	}

	// get first item's url
	item := bow.Find("tr.bookRow")
	URL, _ := item.Find("h3[itemprop='name']").Find("a").Attr("href")
	name := item.Find("h3[itemprop='name']").Find("a").Eq(0).Text()
	author := item.Find("a[itemprop='author']").Eq(0).Text()
	absoluteURL, _ := bow.ResolveStringUrl(URL)

	// exit in unknown case
	if name == "" {
		log.Println("Unknown error occured ", absoluteURL)
		ret.Error = "Unknown error occured " + absoluteURL + bow.Title() + bow.Body()
		return ret
	}
	log.Println("Book: ", name)
	log.Println("Author: ", author)

	// open the book's page
	bow.Open(absoluteURL)
	log.Println("Visiting: ", absoluteURL)

	// get download link
	downloadLink, bool := bow.Find("a.addDownloadedBook").Attr("href")
	if bool == false {
		log.Fatal("Failed to get download link from ", bow.Url())
	}
	log.Println("Downloading from ", downloadLink)

	// redirect to download
	bow.Click("a.addDownloadedBook")

	if bow.Find(".download-limits-error__header").Text() == "Daily limit reached" {
		log.Print("Daily limit reached")
		ret.Error = ("Daily limit reached")
		return ret
	}

	// prepare file to be written
	file, path := prepareFile(name, author)

	// download the book
	_, err = bow.Download(file)
	if err != nil {
		log.Println("Failed to download ", err)
		ret.Error = "Failed to download "
		return ret
	}
	log.Println("Downloaded!")

	defer file.Close()

	// link := uploadFile(filepath)

	ret.Name = name
	ret.URL = upload.UploadFilePath(path)
	ret.Error = "0"
	ret.IP = myip.GetIP()
	return ret
}
