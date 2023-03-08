package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
)

type fileData struct {
	Filename  string
	Timestamp time.Time
	Size      float64
}

func downloadOne(fileData *fileData, srcUrl string, dstPath string) error {

	// Build fileName from fullPath
	u, err := url.Parse(srcUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	u.Path = path.Join(u.Path, fileData.Filename)
	fullURLFile := u.String()
	fileName := filepath.Join(dstPath, fileData.Filename)

	// File exists
	if _, err := os.Stat(fileName); err == nil {
		ext := filepath.Ext(fileName)
		if strings.ToLower(ext) == ".zip" {
			// Valid zip file
			archive, err := zip.OpenReader(fileName)
			if err == nil {
				fmt.Printf("%s exists\n", fileData.Filename)
				defer archive.Close()
				return nil
			}
		} else if strings.ToLower(ext) == ".xlsx" {
			f, err := excelize.OpenFile(fileName)
			if err == nil {
				defer f.Close()
				return nil
			}
		}
	}

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = os.Chtimes(fileName, fileData.Timestamp, fileData.Timestamp)
	if err != nil {
		log.Fatal(err)
		return err
	}

	tr := &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
		Timeout:   60 * time.Second,
		Transport: tr,
	}

	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer file.Close()

	fmt.Printf("Downloaded %s with size %d\n", fullURLFile, size)
	return nil

}

func parseTime(timeStr string) time.Time {
	myDate, err := time.Parse("2006/01/02 15:04", timeStr)
	if err != nil {
		panic(err)
	}
	return myDate
}

func parseSize(sizeStr string) float64 {
	sizes := strings.Split(sizeStr, " ")
	if len(sizes) != 2 || sizes[1] != "KB" {
		panic(errors.New("sizeStr does not contains KB"))
	}
	s := sizes[0]
	s = strings.Replace(s, ",", ".", 1)
	ss, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return ss
}

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

	fileTable := getFileTable(*srcUrl)
	if len(fileTable) == 0 {
		log.Fatal(errors.New("no files in " + *srcUrl))
		os.Exit(1)
	}
	done := make(chan bool)
	go downloadFiles(*srcUrl, dstFullpath, fileTable, done)
	<-done
}

func getFileTable(srcUrl string) []*fileData {
	c := colly.NewCollector(
		colly.AllowedDomains("www.3gpp.org"),
	)
	// Find and download all links
	fileTable := make([]*fileData, 0, 2048)
	c.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			fileData := &fileData{
				Filename:  el.ChildText("td:nth-child(2)"),
				Timestamp: parseTime(el.ChildText("td:nth-child(3)")),
				Size:      parseSize(el.ChildText("td:nth-child(4)")),
			}
			fileTable = append(fileTable, fileData)
		})
	})
	c.Visit(srcUrl)
	return fileTable
}

func downloadFiles(srcUrl string, dstFullpath string, fileTable []*fileData, done chan bool) {
	for _, fileData := range fileTable {
		numTry := 0
		for {
			err := downloadOne(fileData, srcUrl, dstFullpath)
			if err == nil {
				break
			}
			time.Sleep(5 * time.Second)
			numTry = numTry + 1
			if numTry == 5 {
				fmt.Printf("Tried to downloaded file %s for %d times but failed, skipped\n", fileData.Filename, numTry)
				break
			}
		}
	}
	done <- true
	fmt.Println(len(fileTable))
}
