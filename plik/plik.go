package plik

import (
	"github.com/root-gg/plik/plik"
	"log"
	"os"
)

var SERVER string = "https://plik.root.gg"

func UploadFile(path string) string {
	client := plik.NewClient(SERVER)

	upload := client.NewUpload()
	file, err := upload.AddFileFromPath(path)
	// fmt.Println("Uploading ", path)

	err = file.Upload()
	if err != nil {
		log.Println("Failed to upload", err)
	}
	fileURL, err := file.GetURL()
	// fmt.Println("Uploaded!")
	// fmt.Println(fileURL)

	return fileURL.String()
}

func main() {
	f := os.Args[1]
	// log.Println(f[1])
	// if f[1] == "" {
	// log.Println("NO INPUT GIVEN")
	// }
	URL := UploadFile(f)
	log.Printf("LINK: %s", URL)
}
