package main

import "fmt"

func GetNumber() (num int) {
	for {
		fmt.Printf("Enter a number in range [%d,%d]:\n> ", -MaxNumber, MaxNumber)
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Error: Enter a valid integer number.")
			var buf string
			fmt.Scanln(&buf)
			continue
		}

		if num > MaxNumber {
			fmt.Printf("Error: Incorrect number ( %d > %d ).\n", num, MaxNumber)
		} else if num < -MaxNumber {
			fmt.Printf("Error: Incorrect number ( %d < %d ).\n", num, -MaxNumber)
		} else {
			return
		}
	}
}
