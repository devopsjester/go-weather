package location

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
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

func GetZipcode() string {
	zipcodePtr := flag.String("zipcode", "", "A five digit zipcode")
	flag.Parse()

	if len(*zipcodePtr) != 5 {
		fmt.Println("Provide a 5-digit zipcode, please.")
		os.Exit(1)
	}

	return *zipcodePtr
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
