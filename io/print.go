package io

import "fmt"

func PrintResult(resultNum int, textResultNum string, errorCode error) {
	if errorCode == nil {
		fmt.Printf("Result number: %d\n", resultNum)
		fmt.Printf("Number translated into text: %s\n", textResultNum)
	} else {
		fmt.Println(errorCode)
	}
}
