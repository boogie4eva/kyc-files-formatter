package main

import (
	"bitbucket.org/IyenemiTyger/kyc-base64-converter/work"
	"flag"
	"log"
	"sync"
)

var inputDirFlag string
var outputDirFlag string
var numOfWorkers = 10

var wg sync.WaitGroup

func main() {
	flag.StringVar(&inputDirFlag, "sourceDir", inputDirFlag, "-sourceDir=/home/${USER}")
	flag.StringVar(&outputDirFlag, "destDir", "", "destDir=/home/output")
	flag.IntVar(&numOfWorkers, "numOfWorkers", 0, "-numOfWorkers=10")
	flag.Parse()
	if len(inputDirFlag) == 0 {
		log.Fatal("No source directory specified!.")
	}
	if len(outputDirFlag) == 0 {
		log.Printf("Output directory not specified using %s as output .", outputDirFlag)
	}
	if numOfWorkers == 0 {
		log.Printf("Number of workers not specified using the default of %d ", numOfWorkers)
	}
	files, e := work.ReadFromDir(inputDirFlag)
	if e != nil {
		log.Fatalf("Unable to read from input directory %e", e)
	}

	pool := work.New(numOfWorkers, inputDirFlag, outputDirFlag)
	wg.Add(len(files))
	log.Printf("Reading xml files from input  %s", inputDirFlag)
	for file := range files {

	}

}
