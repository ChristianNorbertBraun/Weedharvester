package weedharvester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &assign)

	return assign
}

func (m *master) Find() (url string) {
	return "http://docker:8080"
}
