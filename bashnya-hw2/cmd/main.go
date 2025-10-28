package main

import (
	"bashnya-hw2/constants"
	"bashnya-hw2/convertion"
	"bashnya-hw2/io"
	"bashnya-hw2/text"
	"errors"
)

func main() {
	num := io.GetNumber()

	var resultNum int
	var errorCode error
	resultNum, errorCode = convertion.NumberConversion(num)

	var textResultNum string
	if !errors.Is(errorCode, constants.ErrorService) {
		textResultNum, errorCode = text.NumToText(resultNum)
	}

	io.PrintResult(resultNum, textResultNum, errorCode)
}
