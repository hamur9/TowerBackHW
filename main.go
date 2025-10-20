package main

func main() {
	num := GetNumber()

	resultNum, errorCode := NumberConversion(num)

	textResultNum := NumToText(resultNum)

	PrintResult(resultNum, textResultNum, errorCode)
}
