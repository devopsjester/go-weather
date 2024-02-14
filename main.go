package main

import (
	"fmt"
	"os"
	"go-weather/pkg/location"
	"go-weather/pkg/weather"
)

const OpenWeatherAPIKey = "a015eabe192553962a4cbdb9e7480e45" // Replace with your actual API key


func main() {
	zipcode := location.GetZipcode()

	location, err := location.GetLocation(zipcode)
	if err != nil {
		fmt.Println("Error getting location:", err)
		os.Exit(1)
	}

	temp, err := weather.GetTemperature(location.Places[0].Latitude, location.Places[0].Longitude, OpenWeatherAPIKey)
	if err != nil {
		fmt.Println("Error getting temperature:", err)
		os.Exit(1)
	}

	fmt.Println("The temperature in", location.Places[0].City, location.Places[0].State, "is", temp, "degrees Fahrenheit.")
}