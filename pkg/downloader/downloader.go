package downloader

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/joshuazhu78/gpp-bot/pkg/parser"
)

func downloadOne(fileData *parser.FileData, srcUrl string, dstPath string) error {

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
			err := parser.ParseTdocList(fileName)
			if err != nil {
				return err
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

func DownloadFiles(srcUrl string, dstFullpath string, fileTable []*parser.FileData, done chan bool) {
	for _, fileData := range fileTable {
		err := downloadOne(fileData, srcUrl, dstFullpath)
		if err != nil {
			log.Fatal(errors.New("downloadOne error"))
		}
	}
	done <- true
	fmt.Println(len(fileTable))
}
