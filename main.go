package main

import (
	"encoding/xml"
	"github.com/finkf/gocropy"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: %s <hocr> <dir> ", os.Args[0])
	}
	hocr := gocropy.MustReadHOCR(os.Args[1])
	err := hocr.ConvertToHOCR(os.Args[2])
	if err == nil {
		enc := xml.NewEncoder(os.Stdout)
		enc.Indent(" ", " ")
		enc.Encode(hocr)
		os.Stdout.WriteString("\n")
	} else {
		log.Fatal(err)
	}
}
