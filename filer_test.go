package weedharvester

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestCreateFiler(t *testing.T) {
	filer := NewFiler("http://docker:8888")
	now := time.Now().UTC()
	timeAsString := now.Format(time.RFC3339Nano)
	err := createFile(timeAsString, "test/path", "only a test", &filer)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestReadFiler(t *testing.T) {
	filer := NewFiler("http://docker:8888")
	content := "Only a test"
	err := createFile("test", "test/read", content, &filer)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	reader := filer.Read("test", "test/read")
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if string(data) != content {
		t.Errorf("Expected %s but got %s", content, string(data))
	}

}

func TestReadDirectory(t *testing.T) {
	filer := NewFiler("http://docker:8888")
	content := "Only a test"
	err := createFile("test", "test/path", content, &filer)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	directory := filer.ReadDirectory("test", "")

	if directory.Directory != "test" {
		t.Errorf("Error: Retruned directory is not named test")
	}
}

func createFile(filename string, filepath string, content string, filer *Filer) error {
	data := bytes.NewReader([]byte(content))
	return filer.Create(data, filename, filepath)
}
