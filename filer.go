package weedharvester

import (
	"bytes"
	"fmt"
	"io"
)

// Filer represents a filer for SeaweedFS
type Filer struct {
	url string
}

// Create creates a new file with the given content name and under the given path
func (f *Filer) Create(content io.Reader, filename string, path string) error {
	var b bytes.Buffer
	writer, err := createMultipartForm(&content, &b, "")

	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s/%s", f.url, path, filename)
	return sendMultipartFormData(writer, &b, url)
}
