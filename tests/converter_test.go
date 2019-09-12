package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

const fileSource = "/home/ityger/Projects/Vidicon/sample-ncc"
const outputDir = "base64-converter-output"

func TestDirReading(t *testing.T) {

	t.Log("testing reading of directory with xml files ")
	startTime := time.Now()
	files, e := ioutil.ReadDir(fileSource)
	if e != nil {
		t.Fatalf("Error while reading files %s", e)
	}

	for i, file := range files {
		e := processFile(file)
		if e != nil {
			t.Logf("Error processing file %e", e)
		} else {
			t.Logf("Processing xml file  %d Name: %+v", i, file.Name())
		}

	}
	elapsed := time.Since(startTime)
	t.Logf("Processing time took %s", elapsed)

}

func processFile(info os.FileInfo) error {

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.Mkdir(outputDir, 777)
	}
	filePath := []string{fileSource, info.Name()}
	fileContents, e := ioutil.ReadFile(strings.Join(filePath, "/"))
	if e != nil {
		log.Printf("Unable to process %s file %e", info.Name(), e)
	}

	output := bytes.Replace(fileContents, []byte("&#13;"), []byte(nil), -1)
	dir, e := os.Getwd()
	if e != nil {
		return e
	}
	conversionPath := fmt.Sprintf("%s/%s", dir, info.Name())
	log.Printf("Conversion %s", conversionPath)
	e = ioutil.WriteFile(conversionPath, output, 777)
	if e != nil {
		return e
	}

	return nil

}
