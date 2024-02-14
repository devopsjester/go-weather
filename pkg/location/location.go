package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Country string `json:"country"`
	Places  []struct {
		City      string `json:"place name"`
		State     string `json:"state"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"places"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type LocationService struct {
	Client HttpClient
}

func (ls *LocationService) GetLocation(zipcode string) (*Location, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://api.zippopotam.us/us/%s", zipcode), nil)
	resp, err := ls.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	return &location, nil
}
