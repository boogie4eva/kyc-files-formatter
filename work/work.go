package work

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	inputDir  string
	outputDir string
)

type (
	// Worker must be implemented by types that want to use
	// the work pool.
	Worker interface {
		Task()
	}

	// Pool provides a pool of goroutines that can execute any Worker
	// tasks that are submitted.
	Pool struct {
		work chan Worker
		wg   sync.WaitGroup
	}

	KYCFile struct {
		File os.FileInfo
	}
)

// New creates a new work pool.
func New(maxGoroutines int, sourceDir string, destDir string) *Pool {
	inputDir = sourceDir
	outputDir = destDir
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run submits work to the pool.
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown waits for all the goroutines to shutdown.
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

// KYC file implements Worker interface
func (kyc *KYCFile) Task() {
	log.Printf("Processing file %s", kyc.File.Name())
	err := processFile(kyc.File)
	if err != nil {
		log.Print(fmt.Errorf("unable to process file %e", err))
	}

}

/**
Reads from the specified directory
*/
func ReadFromDir(dir string) ([]os.FileInfo, error) {
	files, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}
	return files, nil
}

// Processes the file passed
func processFile(info os.FileInfo) error {
	dir, e := os.Getwd()
	if e != nil {
		return e
	}
	filePath := []string{inputDir, info.Name()}
	fileContents, e := ioutil.ReadFile(strings.Join(filePath, "/"))
	if e != nil {
		log.Printf("Unable to process %s file %e", info.Name(), e)
	}

	outputForBadCharacter := bytes.Replace(fileContents, []byte("&#13;"), []byte(nil), -1)
	output := bytes.Replace(outputForBadCharacter, []byte("Re-KYC"), []byte(nil), -1)

	//Write the file with thee check contents stripped
	conversionPath := fmt.Sprintf(filepath.Join(dir, outputDir)+"/%s", info.Name())

	e = ioutil.WriteFile(conversionPath, output, 0777)
	if e != nil {
		return e
	}
	log.Printf("Processed file %s", conversionPath)
	return nil

}
