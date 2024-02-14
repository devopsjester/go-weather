package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const OpenWeatherAPIKey = "a015eabe192553962a4cbdb9e7480e45" // Replace with your actual API key

type Place struct {
	City      string `json:"place name"`
	State     string `json:"state"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Location struct {
	Country string   `json:"country"`
	Places []Place `json:"places"`
}

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getTemperature(lat, lon, apiKey string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=imperial", lat, lon, apiKey))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return 0, err
	}

	return weather.Main.Temp, nil
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

	temp, err := getTemperature(location.Places[0].Latitude, location.Places[0].Longitude, OpenWeatherAPIKey)
	if err != nil {
		fmt.Println("Error getting temperature:", err)
		os.Exit(1)
	}

	fmt.Println("The temperature in", location.Places[0].City, location.Places[0].State, "is", temp, "degrees Fahrenheit.")
}