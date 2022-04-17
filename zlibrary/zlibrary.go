package zlibrary

import (
	// "example.com/headless"
	"fmt"
	"strconv"

	"github.com/root-gg/plik/plik"
	"gopkg.in/headzoo/surf.v1"

	// "log"
	"os"
	"strings"
)

type Book struct {
	Name      string
	Author    string
	FileName  string
	URL       string
	UploadURL string
	Size      string
}

type Output struct {
	Book
	Error string
}

var (
	book Book
	O    Output

	Path     = "./download/"
	FilePath string

	FileName string
	Filters  = "?extensions[]=epub"
	// Fallback = "https://u1lib.org/"
	Fallback = "https://1lib.in/"
)

func prepareFile(b Book) *os.File {
	FileName = b.FileName
	FilePath = Path + FileName

	// check if Path exists
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		fmt.Println(Path, " not found, making it.")
		os.Mkdir(Path, 0755)
	}
	// check if file exists
	if _, err := os.Stat(FilePath); !os.IsNotExist(err) {
		fmt.Println("\nFile already been downloaded at ", FilePath)
	}
	//create file
	file, err := os.Create(FilePath)
	if err != nil {
		fmt.Println("Failed to create file \n", err)
		O.Error = "Failed to create file"
		fmt.Println("Folder: ", Path)
	}
	//set book path
	book.FileName = FileName

	return file
}

func downloadBook(url string) {

	B := surf.NewBrowser()
	// open the book's page
	B.Open(url)
	// fmt.Println("Visiting: ", url)

	// store book info
	book.Name = strings.TrimSpace(B.Find(".col-sm-9 > h1:nth-child(1)").Text())
	book.Author = B.Find(".col-sm-9 > i:nth-child(2) > a:nth-child(1)").Eq(0).Text()
	book.URL, _ = B.ResolveStringUrl(url)
	book.FileName = book.Author + " - " + book.Name + ".epub"

	// panic if greater size
	format := B.Find("div.bookDetailsBox").Find("div.property__file").Find(".property_value").Eq(0).Text()
	ext := strings.Split(format, ",")
	unit := strings.Split(ext[1], " ")
	size, _ := strconv.ParseFloat(unit[1], 32)
	if unit[2] == "MB" {
		if size > 8 {
			O.Error = "Larger file"
			panic("Larger file")
		}
	}

	book.Size = fmt.Sprintf("%.2f", size) + unit[2]

	//print book info
	fmt.Println("\nName: ", book.Name)
	fmt.Println("Author: ", book.Author)
	fmt.Println("Size :", book.Size)
	fmt.Println("URL:", book.URL)

	// get download link
	_, bool := B.Find("a.addDownloadedBook").Attr("href")
	if bool == false {
		fmt.Println("Failed to get download link from ", B.Url())
	}

	// fmt.Println("\nDownloading from ", B.Url)

	// redirect to download
	B.Click("a.addDownloadedBook")

	if B.Find(".download-limits-error__header").Text() == "Daily limit reached" {
		fmt.Print("Daily limit reached")
		O.Error = ("Daily limit reached")
		panic(O.Error)
	}

	// prepare file to be written
	file := prepareFile(book)

	// download the book
	_, err := B.Download(file)
	if err != nil {
		fmt.Println("Failed to download ", err)
		O.Error = "Failed to download "
	}
	fmt.Println("Downloaded!")

	defer file.Close()
}

func uploadBook() string {
	client := plik.NewClient("https://plik.root.gg")

	upload := client.NewUpload()
	file, err := upload.AddFileFromPath(FilePath)
	fmt.Println("Uploading ", FilePath)

	err = file.Upload()
	if err != nil {
		fmt.Println("Failed to upload", err)
	}
	fileURL, err := file.GetURL()
	fmt.Println("Uploaded!")
	fmt.Println(fileURL)

	return fileURL.String()
}

func Query(query string) Output {
	if query == "" {
		fmt.Println("Search for something!")
		O.Error = "Search for something!"
		return O
	}

	// base := headless.GetRedirectURL("https://1lib.domains/?redirectUrl=/")
	// if base == "https://1lib.domains/?redirectUrl=/" {
	base := Fallback
	// }
	URL := base + "s/" + query + Filters
	fmt.Println("Querying ", query)

	B := surf.NewBrowser()
	url := URL

	// open search url
	fmt.Println("Visiting ", URL)
	err := B.Open(URL)

	if err != nil {
		fmt.Println("Failed to open link", url, err)
		O.Error = "Failed to open link" + url
		return O
	}
	// exit if no result
	if B.Find("div.notFound").Text() == "On your request nothing has been found. Do you want to create a ZAlert on this query?" {
		fmt.Println("No books found for at ", URL)
		O.Error = "No books found for"
		return O
	}
	// get first item's url
	url, _ = B.Find("div.resItemBox:nth-child(2) > div:nth-child(1) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(1) > h3:nth-child(1) > a:nth-child(1)").Attr("href")
	name := B.Find("div.resItemBox:nth-child(2) > div:nth-child(1) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(1) > h3:nth-child(1) > a:nth-child(1)").Text()

	// exit in unknown case
	if name == "" {
		fmt.Println("Unknown error occured ", url)
		O.Error = "Unknown error occured " + url + B.Title() + B.Body()
		return O
	}

	//download
	book.URL, _ = B.ResolveStringUrl(url)
	downloadBook(book.URL)

	//upload
	book.UploadURL = uploadBook()

	//return
	O.Book = book
	return O
}
