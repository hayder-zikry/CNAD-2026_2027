package main

import "fmt"

func main() {
	libFines()
}

func libFines() {
	var userInput float64

	fmt.Print("Enter number of days overdue:")
	fmt.Scan(&userInput)

	var totalFines float64
	if userInput <= 7 {
		// 3(a)
		totalFines = userInput * 0.1
	} else {
		// 3(b)
		totalFines = userInput * 0.2
	}

	fmt.Printf("%.2f", totalFines)
}
