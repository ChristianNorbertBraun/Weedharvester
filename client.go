package weedharvester

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Report represents all data for
type Report struct {
	FileID  string
	Success bool
}

// Client represents a SeaweedFS client
type Client struct {
	master master
}

// Read reads file with a given fileId
func (c *Client) Read(fileID string) io.Reader {
	location := c.master.Find(fileID)
	resp, err := http.Get(location.PublicURL + "/" + fileID)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode >= 300 {
		panic(fmt.Sprintln("Bad status code reading"))
	}

	return resp.Body
}

// Create creates a file for the given content within the SeaweedFS
func (c *Client) Create(content io.Reader) (string, error) {
	var b bytes.Buffer
	assign := c.master.Assign()

	writer, err := createMultipartForm(&content, &b)

	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s/%s", assign.PublicURL, assign.Fid)

	return assign.Fid, sendMultipartFormData(writer, &b, url)
}

// Delete deletes the file with the given fileId
func (c *Client) Delete() {

}
