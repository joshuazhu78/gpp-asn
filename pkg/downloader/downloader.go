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
	"sync"

	"github.com/joshuazhu78/gpp-bot/pkg/parser"
)

type Downloader struct {
	mu     sync.Mutex
	client *http.Client
}

func NewDownloader() *Downloader {
	return &Downloader{
		client: nil,
	}
}

func (d *Downloader) GetClient() *http.Client {
	d.mu.Lock()
	if d.client == nil {
		d.client = &http.Client{}
	}
	d.mu.Unlock()
	return d.client
}

func (d *Downloader) DownloadOne(fileData *parser.FileData, srcUrl string, dstPath string) (*parser.TDocList, error) {

	// Build fileName from fullPath
	u, err := url.Parse(srcUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	u.Path = path.Join(u.Path, fileData.Filename)
	fullURLFile := u.String()
	fileName := filepath.Join(dstPath, fileData.Filename)
	var tdocList *parser.TDocList = nil

	// File exists
	if _, err := os.Stat(fileName); err == nil {
		ext := filepath.Ext(fileName)
		if strings.ToLower(ext) == ".zip" {
			// Valid zip file
			archive, err := zip.OpenReader(fileName)
			if err == nil {
				fmt.Printf("%s exists\n", fileData.Filename)
				defer archive.Close()
				return nil, nil
			}
		} else if strings.ToLower(ext) == ".xlsx" {
			tdocList, err = parser.ParseTdocList(fileName)
			if err != nil {
				return nil, err
			}
			return tdocList, nil
		}
	}

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return tdocList, err
	}

	err = os.Chtimes(fileName, fileData.Timestamp, fileData.Timestamp)
	if err != nil {
		log.Fatal(err)
		return tdocList, err
	}

	// Put content on file
	resp, err := d.GetClient().Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
		return tdocList, err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
		return tdocList, err
	}

	defer file.Close()

	fmt.Printf("Downloaded %s with size %d\n", fullURLFile, size)
	return tdocList, nil

}

func linkTdocs(tdocList *parser.TDocList, dstFullpath string) {
	for _, tdocEntry := range tdocList.TableEntries {
		docParentPath := filepath.Dir(dstFullpath)
		agendaItemPath := filepath.Join(docParentPath, tdocEntry["Agenda item"])
		os.MkdirAll(agendaItemPath, os.ModePerm)
		target := filepath.Join(dstFullpath, tdocEntry["TDoc"]+".zip")
		if _, err := os.Stat(target); err != nil {
			continue
		}
		tdocSrc := tdocEntry["Source"]
		tdocSrc = strings.Replace(tdocSrc, " ", "_", -1)
		tdocSrc = strings.Replace(tdocSrc, ",", "", -1)
		symLink := filepath.Join(agendaItemPath, tdocEntry["TDoc"]+"_"+tdocSrc+".zip")
		os.Symlink(target, symLink)
	}
}

func DownloadAndLinkTdocs(srcUrl string, dstFullpath string, fileTable []*parser.FileData, j uint) {
	var tdocList *parser.TDocList = nil
	d := NewDownloader()
	guard := make(chan struct{}, j)
	for _, fileData := range fileTable {
		guard <- struct{}{}
		go func(fileData *parser.FileData) {
			tList, err := d.DownloadOne(fileData, srcUrl, dstFullpath)
			if err != nil {
				log.Fatal(errors.New("downloadOne error"))
			}
			if tdocList == nil && tList != nil {
				tdocList = tList
			}
			<-guard
		}(fileData)
	}
	if tdocList != nil {
		linkTdocs(tdocList, dstFullpath)
	}
	fmt.Println(len(fileTable))
}
