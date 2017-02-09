package weedharvester

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Filer represents a filer for SeaweedFS
type Filer struct {
	url string
}

// Directory represents a directory mapped from the seaweed filer
type Directory struct {
	Directory      string         `json:"Directory"`
	Files          []file         `json:"Files"`
	Subdirectories []subdirectory `json:"Subdirectories"`
}

type subdirectory struct {
	Name string `json:"Name"`
	ID   int    `json:"Id"`
}

type file struct {
	Name string `json:"name"`
	Fid  string `json:"fid"`
}

// Create creates a new file with the given content name and under the given path
func (f *Filer) Create(content io.Reader, filename string, path string) error {
	var b bytes.Buffer
	writer, err := createMultipartForm(&content, &b)

	if err != nil {
		return err
	}

	var url string
	if len(path) == 0 {
		url = fmt.Sprintf("%s%s", addSlashIfNeeded(f.url), filename)
	} else {
		url = fmt.Sprintf("%s%s%s", addSlashIfNeeded(f.url), addSlashIfNeeded(path), filename)
	}

	return sendMultipartFormData(writer, &b, url)
}

func (f *Filer) Read(filename string, path string) (io.Reader, error) {
	var url string
	if len(path) == 0 {
		url = fmt.Sprintf("%s%s", addSlashIfNeeded(f.url), filename)
	} else {
		url = fmt.Sprintf("%s%s%s", addSlashIfNeeded(f.url), addSlashIfNeeded(path), filename)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Bad StatusCode reading: %s/%s", filename, path)
	}

	return resp.Body, nil
}

// ReadDirectory returns all files contained in a given directory
func (f *Filer) ReadDirectory(path string, lastFileName string) (*Directory, error) {
	var url string
	if len(lastFileName) != 0 {
		url = fmt.Sprintf("%s%s?lastFileName=%s", addSlashIfNeeded(f.url), addSlashIfNeeded(path), lastFileName)
	} else {
		url = fmt.Sprintf("%s%s", addSlashIfNeeded(f.url), addSlashIfNeeded(path))
	}
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	directory := Directory{}
	err = decodeJSON(resp.Body, &directory)

	if err != nil {
		return nil, err
	}
	return &directory, err
}
