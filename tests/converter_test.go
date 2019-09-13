package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const fileSource = "/home/ityger/Projects/Vidicon/sample-ncc"
const outputDir = "base64-formatter-output"

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
			t.Errorf("Error processing file %e", e)
		} else {
			t.Logf("Processing xml file  %d Name: %+v", i, file.Name())
		}

	}
	elapsed := time.Since(startTime)
	t.Logf("Processing time took %s", elapsed)

}

func processFile(info os.FileInfo) error {
	dir, e := os.Getwd()
	if e != nil {
		return e
	}
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err := os.Mkdir(filepath.Join(dir, outputDir), 0777)
		if err != nil {
			return err
		}
	}
	filePath := []string{fileSource, info.Name()}
	fileContents, e := ioutil.ReadFile(strings.Join(filePath, "/"))
	if e != nil {
		log.Printf("Unable to process %s file %e", info.Name(), e)
	}

	output := bytes.Replace(fileContents, []byte("&#13;"), []byte(nil), -1)

	//conversionPath := fmt.Sprintf("%s/%s/%s", dir, outputDir, info.Name())
	conversionPath := fmt.Sprintf(filepath.Join(dir, outputDir)+"/%s", info.Name())
	log.Printf("Conversion %s", conversionPath)
	e = ioutil.WriteFile(conversionPath, output, 0777)
	if e != nil {
		return e
	}

	return nil

}
