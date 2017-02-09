package weedharvester

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Client represents a SeaweedFS client
type Client struct {
	master master
}

// Read reads file with a given fileId
func (c *Client) Read(fileID string) (io.Reader, error) {
	location, err := c.master.Find(fileID)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(location.PublicURL + "/" + fileID)

	if err != nil {
		log.Printf("Error while sending get to %s/%s", location.PublicURL, fileID)
		return nil, err
	}

	if resp.StatusCode >= 300 {
		log.Printf("Status %d while reading from %s/%s", resp.StatusCode, location.PublicURL, fileID)
		return nil, errors.New("Bad StatusCode")
	}

	return resp.Body, nil
}

// Create creates a file for the given content within the SeaweedFS
func (c *Client) Create(content io.Reader) (string, error) {
	var b bytes.Buffer
	assign, err := c.master.Assign()

	if err != nil {
		return "", err
	}

	writer, err := createMultipartForm(&content, &b)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s", assign.PublicURL, assign.Fid)

	return assign.Fid, sendMultipartFormData(writer, &b, url)
}

// Delete deletes the file with the given fileId
func (c *Client) Delete() {

}
