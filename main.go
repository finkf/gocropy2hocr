package main

import (
	"encoding/xml"
	//	"fmt"
	"github.com/finkf/gocropy"
	"log"
	"os"
)

func main() {
	// llocs, err := gocropy.ReadLlocs(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }
	// for i := range llocs {
	// 	fmt.Printf("%s\n", llocs[i])
	// }
	if len(os.Args) != 3 {
		log.Fatal("Usage: %s <hocr> <dir> ", os.Args[0])
	}
	hocr := gocropy.MustReadHocr(os.Args[1])
	err := hocr.ConvertToHocr(os.Args[2])
	if err == nil {
		enc := xml.NewEncoder(os.Stdout)
		enc.Indent(" ", " ")
		enc.Encode(hocr)
	} else {
		log.Fatal(err)
	}
}
