package main

import (
	"flag"
	"fmt"
)

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

func main() {
	zipcode := getZipcode()

	// Print the zipcode
	fmt.Println("The zipcode you entered is:", zipcode)
}