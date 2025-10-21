package io

import (
	"bashnya-hw2/constants"
	"fmt"
)

func GetNumber() (num int) {
	for {
		fmt.Printf("Enter a number in range [%d,%d]:\n> ", -constants.MaxNumber, constants.MaxNumber)
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Error: Enter a valid integer number.")
			var buf string
			fmt.Scanln(&buf)
			continue
		}

		if num > constants.MaxNumber {
			fmt.Printf("Error: Incorrect number ( %d > %d ).\n", num, constants.MaxNumber)
		} else if num < -constants.MaxNumber {
			fmt.Printf("Error: Incorrect number ( %d < %d ).\n", num, -constants.MaxNumber)
		} else {
			return
		}
	}
}
