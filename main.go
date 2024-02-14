package main

import (
	"fmt"
	"os"
)

const OpenWeatherAPIKey = "a015eabe192553962a4cbdb9e7480e45" // Replace with your actual API key


func main() {
	zipcode := GetZipcode()

	location, err := GetLocation(zipcode)
	if err != nil {
		fmt.Println("Error getting location:", err)
		os.Exit(1)
	}

	temp, err := GetTemperature(location.Places[0].Latitude, location.Places[0].Longitude, OpenWeatherAPIKey)
	if err != nil {
		fmt.Println("Error getting temperature:", err)
		os.Exit(1)
	}

	fmt.Println("The temperature in", location.Places[0].City, location.Places[0].State, "is", temp, "degrees Fahrenheit.")
}