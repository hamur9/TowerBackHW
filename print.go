package main

import "fmt"

func PrintResult(resultNum int, textResultNum string, errorCode string) {
	if errorCode == ErrorOK {
		fmt.Printf("Result number: %d\n", resultNum)
		fmt.Printf("Number translated into text: %s\n", textResultNum)
	} else {
		fmt.Println(errorCode)
	}
}
