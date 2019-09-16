package main

import (
	"bitbucket.org/IyenemiTyger/kyc-base64-converter/work"
	"bytes"
	"flag"
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

	outputForBadCharacter := bytes.Replace(fileContents, []byte("&#13;"), []byte(nil), -1)
	output := bytes.Replace(outputForBadCharacter, []byte("Re-KYC"), []byte(nil), -1)

	//Write the file with thee check contents stripped
	conversionPath := fmt.Sprintf(filepath.Join(dir, outputDir)+"/%s", info.Name())
	log.Printf("Conversion %s", conversionPath)
	e = ioutil.WriteFile(conversionPath, output, 0777)
	if e != nil {
		return e
	}

	return nil

}

func TestWorkSynchronization(t *testing.T) {

	var inputDirFlag string
	var outputDirFlag string
	var numOfWorkers = 10

	flag.StringVar(&inputDirFlag, "sourceDir", inputDirFlag, "-sourceDir=/home/${USER}")
	flag.StringVar(&outputDirFlag, "destDir", "", "destDir=/home/output")
	flag.IntVar(&numOfWorkers, "numOfWorkers", 0, "-numOfWorkers=10")
	flag.Parse()
	if len(inputDirFlag) == 0 {
		t.Fatal("No source directory specified!.")
	}
	if len(outputDirFlag) == 0 {
		t.Logf("Output directory not specified using %s as output .", outputDir)
	}
	if numOfWorkers == 0 {
		t.Logf("Number of workers not specified using the default of %d ", numOfWorkers)
	}
	files, e := ReadFromDir(inputDirFlag)
	if e != nil {
		t.Fatalf("Unable to read from input directory %e", e)

	}

	pool := work.New(numOfWorkers)

	wg.Add(100 * len(names))
	t.Logf("Reading xml files from input  %s", inputDirFlag)
	for file := range files {

	}
}

/**
Reads from the specified directory
*/
func ReadFromDir(dir string) ([]os.FileInfo, error) {
	files, e := ioutil.ReadDir(fileSource)
	if e != nil {
		return nil, e
	}
	return files, nil
}
