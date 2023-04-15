package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joshuazhu78/gpp-asn/pkg/parser"
)

func main() {
	specTxt := flag.String("specTxt", "38331-gb0.txt", "spec text file with .txt extension and UTF-8 format")
	asnPath := flag.String("asnPath", "./asn/", "destination path to save the asn1 syntex")

	flag.Parse()
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()

	err := parser.ParseSpec(*specTxt, *asnPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
