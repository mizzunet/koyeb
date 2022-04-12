package upload

import (
	"github.com/root-gg/plik/plik"
	"log"
)

func UploadFilePath(f string) string {
	client := plik.NewClient("https://plik.root.gg")

	upload := client.NewUpload()
	file, err := upload.AddFileFromPath(f)
	log.Println("Uploading.. ", f)

	err = file.Upload()
	if err != nil {
		log.Println("some erre", err)
	}
	uploadURL, err := upload.GetURL()
	fileURL, err := file.GetURL()
	log.Println("Uploaded")
	log.Println(uploadURL)

	return fileURL.String()
}
