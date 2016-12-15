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

type volume struct {
	VolumeID  string     `json:"volumeId"`
	Locations []location `json:"locations"`
}

type location struct {
	URL       string `json:"url"`
	PublicURL string `json:"publicUrl"`
}

func (m *master) Assign() assignment {
	completeURL := m.url + "/dir/assign"
	resp, err := http.Get(completeURL)
	if err != nil {
		panic(fmt.Sprintf("Error: Unable to ask for assignment at: %s", completeURL))
	}

	if resp.StatusCode >= 300 {
		panic(fmt.Sprintln("Received bad statuscode at assignment"))
	}
	assign := assignment{}
	err = decodeJSON(resp.Body, &assign)

	return assign
}

func (m *master) Find(fileID string) location {
	completeURL := m.url + "/dir/lookup?volumeId=" + fileID
	resp, err := http.Get(completeURL)
	if err != nil {
		panic(fmt.Sprintf("Error: Unable to lookup for volume at: %s", completeURL))
	}

	if resp.StatusCode >= 300 {
		panic(fmt.Sprintln("Received bad statuscode at lookup"))
	}
	volume := volume{}
	err = decodeJSON(resp.Body, &volume)

	if err != nil {
		panic(fmt.Sprintf("Error: Unable to parse response from %s", completeURL))
	}

	return volume.Locations[0]
}
