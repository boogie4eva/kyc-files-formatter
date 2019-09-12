package main

import (
	"io/ioutil"
	"testing"
)

func TestDirReading(t *testing.T) {

	t.Log("testing reading of directory with xml files ")
	files, e := ioutil.ReadDir("/home/ityger/Projects/Vidicon/sample-ncc")
	if e != nil {
		t.Fatalf("Error while reading files %s", e)
	}

	for i, file := range files {
		t.Logf("Processing xml file  %d Name: %+v", i, file)
	}
	//t.Logf("All files in specified directory %q", file)

}
