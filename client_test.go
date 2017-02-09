package weedharvester

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestCreate(t *testing.T) {
	client := NewClient(*masterURL)

	data := bytes.NewReader([]byte("Only a test"))
	_, err := client.Create(data)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestRead(t *testing.T) {
	client := NewClient(*masterURL)
	dataString := "Only a test"
	data := bytes.NewReader([]byte(dataString))
	fileID, err := client.Create(data)
	if err != nil {
		t.Errorf("Error:%s", err)
	}

	reader, err := client.Read(fileID)
	if err != nil {
		t.Errorf("Error: Unable to read file: %s %s", fileID, err)
	}
	byteArray, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	responseData := string(byteArray)

	if responseData != dataString {
		t.Errorf("Expected %s but got %s", dataString, responseData)
	}
}
