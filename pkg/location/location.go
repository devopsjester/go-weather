package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Place struct {
	City      string `json:"place name"`
	State     string `json:"state"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Location struct {
	Country string  `json:"country"`
	Places  []Place `json:"places"`
}

func GetLocation(zipcode string) (*Location, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.zippopotam.us/us/%s", zipcode))
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
