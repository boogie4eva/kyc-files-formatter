package main

import (
	"flag"
	"github.com/boogie4eva/kyc-files-formatter/work"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var inputDirFlag string
var outputDirFlag string
var numOfWorkers = 50
var wg sync.WaitGroup

func main() {
	inDirFlag := flag.String("sourceDir", "", "-sourceDir=/home/${USER}")
	outDirFlag := flag.String("destDir", "", "-destDir=output")
	Workers := flag.Int("numOfWorkers", 0, "-numOfWorkers=10")
	flag.Parse()
	inputDirFlag = *inDirFlag
	outputDirFlag = *outDirFlag
	numOfWorkers = *Workers

	if len(inputDirFlag) == 0 {
		log.Fatal("No source directory specified!.")
	}
	if len(outputDirFlag) == 0 {
		outputDirFlag = "base64-file-formatter-output"
		log.Printf("Output directory not specified using %s as output .", outputDirFlag)
	}
	if numOfWorkers == 0 {
		log.Printf("Number of workers not specified using the default of %d ", numOfWorkers)
	}
	// Get the current working directory
	dir, e := os.Getwd()
	if e != nil {
		log.Fatalf("Unable to get the current working directory %e", e)
	}
	if _, err := os.Stat(outputDirFlag); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err := os.Mkdir(filepath.Join(dir, outputDirFlag), 0777)
		if err != nil {
			log.Fatalf("Unable to create the output directory %e", e)
		}
	}
	// Read the files from the specified source
	files, e := work.ReadFromDir(inputDirFlag)
	if e != nil {
		log.Fatalf("Unable to read from input directory %e", e)
	}
	pool := work.New(numOfWorkers, inputDirFlag, outputDirFlag)
	log.Printf("Reading xml files from input  %s", inputDirFlag)
	startTime := time.Now()
	wg.Add(len(files))
	for _, file := range files {
		kycFile := work.KYCFile{File: file}
		go func() {
			pool.Run(&kycFile)
			wg.Done()

		}()
	}
	wg.Wait()
	pool.Shutdown()
	endTime := time.Since(startTime)
	log.Printf("Total execution time %s", endTime)

}
