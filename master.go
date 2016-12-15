package weedharvester

import (
	"fmt"
	"net/http"
)

type master struct {
	url string
}
type assignment struct {
	Count     int    `json:"count"`
	Fid       string `json:"fid"`
	URL       string `json:"url"`
	PublicURL string `json:"publicUrl"`
}

func (m *master) Assign() assignment {
	completeURL := m.url + "/dir/assign"
	response, err := http.Get(completeURL)
	if err != nil {
		panic(fmt.Sprintf("Error: Unable to ask for assignment at: %s", completeURL))
	}

	assign := assignment{}
	decodeJSON(response.Body, &assign)

	return assign
}

func (m *master) Find() (url string) {
	return "http://docker:8080"
}
