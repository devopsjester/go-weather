package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Place struct {
	City  string `json:"place name"`
	State string `json:"state"`
}

type Location struct {
	Country string   `json:"country"`
	Places []Place `json:"places"`
}

// getZipcode function to get the zipcode from command-line arguments
func getZipcode() string {
	zipcodePtr := flag.String("zipcode", "", "A five digit zipcode")
	flag.Parse()

	// Check if zipcode is a 5-digit number
	if len(*zipcodePtr) != 5 {
		fmt.Println("Please provide a 5-digit zipcode.")
		os.Exit(1)
	}

	return *zipcodePtr
}

func getLocation(zipcode string) (*Location, error) {
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

func main() {
	zipcode := getZipcode()

	location, err := getLocation(zipcode)
	if err != nil {
		fmt.Println("Error getting location:", err)
		os.Exit(1)
	}

	fmt.Println("The city is:", location.Places[0].City)
	fmt.Println("The state is:", location.Places[0].State)
}