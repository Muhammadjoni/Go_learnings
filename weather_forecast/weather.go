// Package weather is a comment practicing exercise
package main

import "fmt"

// CurrentTemp = the actual temp provided to be used afterwards inside functions
// CurrentLocation = the current location as the name literally describes
var CurrentTemp int
var CurrentLocation string

// Forecast returns a string value equal to the temperature based on currentTemp of currentLocation in degrees Celsius.
func Forecast(temp int, loc string) string {
	var condition string

	if temp < 0 {
		condition = "cold"
	} else if temp > 0 && temp < 15 {
		condition = "it is moderate and little cool"
	} else {
		condition = "it is terrible"
	}

	return fmt.Sprintf("The current weather condition is %s", condition)
}

func main() {
	fmt.Println(Forecast(14, "Berlin"))
}
