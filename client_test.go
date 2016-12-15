package weedharvester

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestCreate(t *testing.T) {
	client := NewClient("http://docker:9333")

	data := bytes.NewReader([]byte("Only a test"))
	_, err := client.Create(data)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestRead(t *testing.T) {
	client := NewClient("http://docker:9333")
	dataString := "Only a test"
	data := bytes.NewReader([]byte(dataString))
	fileID, err := client.Create(data)
	if err != nil {
		t.Errorf("Error:%s", err)
	}

	reader := client.Read(fileID)
	byteArray, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	responseData := string(byteArray)

	if responseData != dataString {
		t.Errorf("Expected %s but got %s", dataString, responseData)
	}
}
