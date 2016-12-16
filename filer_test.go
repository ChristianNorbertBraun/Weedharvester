package weedharvester

import (
	"bytes"
	"testing"
)

func TestCreateFiler(t *testing.T) {
	filer := NewFiler("http://docker:8888")

	data := bytes.NewReader([]byte("Only a test"))
	err := filer.Create(data, "test", "test/path")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
