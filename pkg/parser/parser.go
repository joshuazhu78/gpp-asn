package parser

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ParseSpec(srcFile, dstPath string) error {
	readFile, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	fileName := filepath.Base(srcFile)
	fileNameNoExt := strings.Split(fileName, ".")
	dstFileName := filepath.Join(dstPath, fileNameNoExt[0]+".asn1")
	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	datawriter := bufio.NewWriter(dstFile)
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	enter := false
	for fileScanner.Scan() {
		if !enter {
			if strings.Contains(fileScanner.Text(), "-- ASN1START") {
				enter = true
				continue
			}
		} else {
			if strings.Contains(fileScanner.Text(), "-- ASN1STOP") {
				enter = false
				continue
			} else {
				_, _ = datawriter.WriteString(fileScanner.Text() + "\n")
			}
		}
	}

	readFile.Close()
	datawriter.Flush()
	dstFile.Close()
	return nil
}
