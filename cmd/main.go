package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/joshuazhu78/gpp-bot/pkg/downloader"
	"github.com/joshuazhu78/gpp-bot/pkg/parser"
)

func main() {
	srcUrl := flag.String("srcUrl", "https://www.3gpp.org/ftp/TSG_RAN/WG1_RL1/TSGR1_112/Docs", "source URL")
	dstDir := flag.String("dstDir", "/home/yzhu/Workspace/3gpp_ftp", "destination dir to save contributions")

	flag.Parse()
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()

	dstSubpath := strings.ReplaceAll(*srcUrl, "https://www.3gpp.org/ftp/", "")
	dstFullpath := path.Join(*dstDir, dstSubpath)
	if err := os.MkdirAll(dstFullpath, os.ModePerm); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileTable := parser.GetFileTable(*srcUrl)
	if len(fileTable) == 0 {
		log.Fatal(errors.New("no files in " + *srcUrl))
		os.Exit(1)
	}
	done := make(chan bool)
	go downloader.DownloadAndLinkTdocs(*srcUrl, dstFullpath, fileTable, done)
	<-done
}
