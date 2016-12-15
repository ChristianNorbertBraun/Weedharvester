package weedharvester

import "testing"
import "bytes"

func TestCreate(t *testing.T) {
	client := NewClient("http://docker:9333")

	data := bytes.NewReader([]byte("Only a test"))
	_, err := client.Create(data)
	if err != nil {
		t.Errorf("Error:%s", err)
	}
}
