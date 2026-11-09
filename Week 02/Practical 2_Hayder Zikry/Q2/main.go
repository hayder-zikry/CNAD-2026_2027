package main

import (
	"fmt"
)

func main() {
	var userInput string

	fmt.Print("Enter a number to convert:")
	fmt.Scan(&userInput)

	var numberToMorseCode = map[rune]string{
		'1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----",
	}

	for _, char := range userInput {
		if code, found := numberToMorseCode[char]; found {
			fmt.Printf("%s   ", code)
		}
	}
}
