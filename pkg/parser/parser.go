package parser

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
)

type FileData struct {
	Filename  string
	Timestamp time.Time
	Size      float64
}

type TableEntry map[string]string
type TDocList struct {
	TableHeader  []string
	TableEntries []TableEntry
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

func GetFileTable(srcUrl string) []*FileData {
	c := colly.NewCollector(
		colly.AllowedDomains("www.3gpp.org"),
	)
	// Find and download all links
	fileTable := make([]*FileData, 0, 2048)
	c.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			fileData := &FileData{
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

func ParseTdocList(fileName string) (*TDocList, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("TDoc_List")
	if err != nil {
		return nil, err
	}
	tdocList := &TDocList{
		TableHeader:  make([]string, 0, 128),
		TableEntries: make([]TableEntry, 0, 1024),
	}
	for i, row := range rows {
		// Table header
		if i == 0 {
			tdocList.TableHeader = append(tdocList.TableHeader, row...)
		} else {
			tableEntry := make(map[string]string)
			for j, colCell := range row {
				tableEntry[tdocList.TableHeader[j]] = colCell
			}
			tdocList.TableEntries = append(tdocList.TableEntries, tableEntry)
		}
	}
	return tdocList, nil
}
