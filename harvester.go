package weedharvester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

// NewClient creates a new client. The given url has to be the address of the master server
func NewClient(url string) Client {
	return Client{master: master{url: url}}
}

// NewFiler creates a new Filer with the given url as the host address
func NewFiler(url string) Filer {
	return Filer{url: url}
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func createMultipartForm(content *io.Reader, b *bytes.Buffer, fieldname string) (*multipart.Writer, error) {
	writer := multipart.NewWriter(b)
	var part io.Writer
	var err error
	if len(fieldname) == 0 {
		part, err = createFormFile(writer, "file", "")
	} else {
		part, err = createFormFile(writer, fieldname, "")
	}
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, *content)
	if err != nil {
		return nil, err
	}

	writer.Close()

	return writer, nil

}

func sendMultipartFormData(writer *multipart.Writer, b *bytes.Buffer, url string) error {
	resp, err := http.Post(
		url,
		writer.FormDataContentType(),
		b)

	if err != nil {
		fmt.Println("Unable to send content!")
		return err
	}

	if resp.StatusCode >= 300 {
		err = fmt.Errorf("bad status: %s", resp.Status)
		return err
	}

	return nil
}

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