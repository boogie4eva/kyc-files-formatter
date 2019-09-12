package main

import (
	"io/ioutil"
	"testing"
)

func TestDirReading(t *testing.T) {

	t.Log("testing reading of directory with xml files ")
	file, e := ioutil.ReadDir("../sample-ncc")
	if e != nil {
		t.Errorf("Error while reading files %s", e)

	}
	t.Logf("All files in specified directory %s", file)

}
