package main

import (
	"fmt"
	"flag"
	"os"
	"go-weather/pkg/location"
	"go-weather/pkg/weather"
)

const OpenWeatherAPIKey = "a015eabe192553962a4cbdb9e7480e45" // Replace with your actual API key

func getZipcodeFromCLIArgument() string {
	zipcodePtr := flag.String("zipcode", "", "A five digit zipcode")
	flag.Parse()

	if len(*zipcodePtr) != 5 {
		fmt.Println("Provide a 5-digit zipcode, please.")
		os.Exit(1)
	}

	return *zipcodePtr
}

func getWeather(zipcode string, getLocation func(string) (*location.Location, error), getTemperature func(string, string, string) (float64, error)) (string, string, float64, error) {
	location, err := getLocation(zipcode)
	if err != nil {
		return "", "", 0, fmt.Errorf("error getting location: %w", err)
	}

	temp, err := getTemperature(location.Places[0].Latitude, location.Places[0].Longitude, OpenWeatherAPIKey)
	if err != nil {
		return "", "", 0, fmt.Errorf("error getting temperature: %w", err)
	}

	return location.Places[0].City, location.Places[0].State, temp, nil
}

func main() {
	zipcode := getZipcodeFromCLIArgument()
	city, state, temp, err := getWeather(zipcode, location.GetLocation, weather.GetTemperature)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("The temperature in", city, state, "is", temp, "degrees Fahrenheit.")
}