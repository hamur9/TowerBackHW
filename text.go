package main

func NumToText(num int) (resultText string) {
	if num == 0 {
		return "ноль"
	}

	lib := numberLibsGet()

	if num >= 1000 {
		hundreds := num / 1000
		n1 := hundreds / 100
		n2 := hundreds / 10 % 10
		n3 := hundreds % 10
		if n1 != 0 {
			resultText += lib.hundreds[n1] + " "
		}

		if hundreds%100 > 10 && hundreds%100 < 20 {
			resultText += lib.teens[hundreds%100] + " "
		} else {
			if n2 != 0 {
				resultText += lib.tens[n2] + " "
			}
			if n3 != 0 {
				if n3 == 1 {
					resultText += "одна" + " "
				} else if n3 == 2 {
					resultText += "две" + " "
				} else {
					resultText += lib.units[n3] + " "
				}
			}
		}
		lastHundredNum := hundreds % 10
		middleHundredNum := hundreds / 10 % 10
		if lastHundredNum == 1 && middleHundredNum != 1 {
			resultText += "тысяча "
		} else if (lastHundredNum == 2 || lastHundredNum == 3 || lastHundredNum == 4) && middleHundredNum != 1 {
			resultText += "тысячи "
		} else {
			resultText += "тысяч "
		}
	}
	n1 := num % 1000 / 100
	n2 := num % 100 / 10
	n3 := num % 10
	if n1 != 0 {
		resultText += lib.hundreds[n1] + " "
	}
	if num%100 > 10 && num%100 < 20 {
		resultText += lib.teens[num%100] + " "
	} else {
		if n2 != 0 {
			resultText += lib.tens[n2] + " "
		}
		if n3 != 0 {
			resultText += lib.units[n3]
		}
	}

	return
}
