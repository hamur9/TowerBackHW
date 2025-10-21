package main

import (
	"bashnya-hw2/convertion"
	"bashnya-hw2/io"
	"bashnya-hw2/text"
)

func main() {
	num := io.GetNumber()

	resultNum, errorCode := convertion.NumberConversion(num)

	textResultNum := text.NumToText(resultNum)

	io.PrintResult(resultNum, textResultNum, errorCode)
}
