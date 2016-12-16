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

// Create creates a new file with the given content name and under the given path
func (f *Filer) Create(content io.Reader, filename string, path string) error {
	var b bytes.Buffer
	writer, err := createMultipartForm(&content, &b)

	if err != nil {
		return err
	}

	var url string
	if len(path) == 0 {
		url = fmt.Sprintf("%s/%s", f.url, filename)
	} else {
		url = fmt.Sprintf("%s/%s/%s", f.url, path, filename)
	}

	return sendMultipartFormData(writer, &b, url)
}

func (f *Filer) Read(filename string, path string) io.Reader {
	var url string
	if len(path) == 0 {
		url = fmt.Sprintf("%s/%s", f.url, filename)
	} else {
		url = fmt.Sprintf("%s/%s/%s", f.url, path, filename)
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode >= 300 {
		panic(fmt.Sprintln("Bad status code reading"))
	}

	return resp.Body
}
