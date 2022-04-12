package main

import (
	"example.com/zlib"
	"log"
	"os"
	"strings"
)

func main() {
	q := os.Args[1:]
	query := strings.Join(q, " ")
	o := zlib.DownloadBook(query)
	log.Println(o)
}
