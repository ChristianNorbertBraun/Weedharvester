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

func (m *master) Assign() (*assignment, error) {
	completeURL := addSlashIfNeeded(m.url) + "dir/assign"
	resp, err := http.Get(completeURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Bad StatusCode reading: %s", completeURL)
	}
	assign := assignment{}
	err = decodeJSON(resp.Body, &assign)

	return &assign, nil
}

func (m *master) Find(fileID string) (*location, error) {
	completeURL := addSlashIfNeeded(m.url) + "dir/lookup?volumeId=" + fileID
	resp, err := http.Get(completeURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Bad StatusCode reading: %s", completeURL)
	}
	volume := volume{}
	err = decodeJSON(resp.Body, &volume)

	if err != nil {
		return nil, err
	}

	return &volume.Locations[0], nil
}
