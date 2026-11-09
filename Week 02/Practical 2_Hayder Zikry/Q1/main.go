package main

import (
	"fmt"
	"math"
)

func main() {
	var userInput float64

	fmt.Print("Enter dewpoint temperature in Celsius: ")
	fmt.Scan(&userInput)

	exponent := (7.5 * userInput) / (237.3 + userInput)
	result := math.Pow(10, float64(exponent))
	pressure := 6.11 * result

	fmt.Printf("The vapor pressure is %.2f millibars", pressure)
}
