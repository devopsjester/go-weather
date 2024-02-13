package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a string flag for zipcode with a default value and a short description.
	zipcodePtr := flag.String("zipcode", "", "A five digit zipcode")

	// Parse the command-line arguments
	flag.Parse()

	// Check if zipcode is a 5-digit number
	if len(*zipcodePtr) != 5 {
		fmt.Println("Please provide a 5-digit zipcode.")
		return
	}

	// Print the zipcode
	fmt.Println("The zipcode you entered is:", *zipcodePtr)
	
}