package weedharvester

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"
)

func TestCreateFiler(t *testing.T) {
	filer := NewFiler(*filerURL)
	now := time.Now().UTC()
	timeAsString := now.Format(time.RFC3339Nano)
	err := createFile(timeAsString, "test/path", "only a test", &filer)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestReadFiler(t *testing.T) {
	filer := NewFiler(*filerURL)
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
	filer := NewFiler(*filerURL)
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

func TestDeleteFile(t *testing.T) {
	filer := NewFiler(*filerURL)
	content := "Only a test"
	err := createFile("test", "test/path/delete", content, &filer)
	err = createFile("test1", "test/path/delete", content, &filer)
	err = createFile("test2", "test/path/delete", content, &filer)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	err = filer.DeleteFilesUntil("test/path/delete", "test1")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	directory, err := filer.ReadDirectory("test/path/delete", "")

	if err != nil {
		t.Errorf("Error: Unable to read directory %s", err)
	}

	if len(directory.Files) != 1 {
		t.Errorf("Error: %s", directory)
	}
}

func createFile(filename string, filepath string, content string, filer *Filer) error {
	data := bytes.NewReader([]byte(content))
	return filer.Create(data, filename, filepath)
}
