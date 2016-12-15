package weedharvester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
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

// NewClient creates a new client. The given url has to be the address of the master server
func NewClient(url string) Client {
	return Client{master: master{url: url}}
}

// Create creates a file for the given content within the SeaweedFS
func (c *Client) Create(content io.Reader) (string, error) {
	assign := c.master.Assign()
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	part, err := createFormFile(writer, "file", "")

	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, content)
	if err != nil {
		return "", err
	}

	writer.Close()
	resp, err := http.Post(
		fmt.Sprintf("%s/%s", assign.PublicURL, assign.Fid),
		writer.FormDataContentType(),
		&b)

	if err != nil {
		fmt.Println("Unable to send content!")
		return "", err
	}

	if resp.StatusCode >= 300 {
		err = fmt.Errorf("bad status: %s", resp.Status)
		return "", err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	return assign.Fid, nil
}

// Read reads file with a given fileId
func (c *Client) Read(fileID string) *io.Reader {

	return nil
}

// Delete deletes the file with the given fileId
func (c *Client) Delete() {

}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func createFormFile(writer *multipart.Writer, fieldname, mime string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"`,
			escapeQuotes(fieldname)))
	if len(mime) == 0 {
		mime = "application/octet-stream"
	}
	h.Set("Content-Type", mime)
	return writer.CreatePart(h)
}
func decodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
