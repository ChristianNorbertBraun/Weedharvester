package weedharvester

import (
	"bytes"
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

	reader, err := filer.Read("test", "test/read")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
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

	directory, err := filer.ReadDirectory("test/path", "")

	if err != nil {
		t.Errorf("Error: Unable to read directory %s", err)
	}

	if directory.Directory != "/test/path/" {
		t.Errorf("Error: Returned directory is not named test but %s", directory.Directory)
	}
}

func createFile(filename string, filepath string, content string, filer *Filer) error {
	data := bytes.NewReader([]byte(content))
	return filer.Create(data, filename, filepath)
}
