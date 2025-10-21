package text

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
		lastTwo := hundreds % 100
		if n1 != 0 {
			resultText += lib.hundreds[n1] + " "
		}

		if lastTwo > 10 && lastTwo < 20 {
			resultText += lib.teens[lastTwo] + " "
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
		middleHundredNum := hundreds / 10 % 10
		if n3 == 1 && middleHundredNum != 1 {
			resultText += "тысяча "
		} else if (n3 == 2 || n3 == 3 || n3 == 4) && middleHundredNum != 1 {
			resultText += "тысячи "
		} else {
			resultText += "тысяч "
		}
	}

	n1 := num % 1000 / 100
	n2 := num % 100 / 10
	n3 := num % 10
	lastTwo := num % 100
	if n1 != 0 {
		resultText += lib.hundreds[n1] + " "
	}
	if lastTwo > 10 && lastTwo < 20 {
		resultText += lib.teens[num%100] + " "
	} else {
		if n2 != 0 {
			resultText += lib.tens[n2] + " "
		}
		if n3 != 0 {
			resultText += lib.units[n3]
		}
	}

	resultText = removeLeadAndEndSpaces(resultText)

	return
}

func removeLeadAndEndSpaces(text string) (result string) {
	if text == "" {
		return ""
	}

	i := len(text) - 1
	for i >= 0 && text[i] == ' ' {
		i--
	}

	k := 0
	for k < len(text) && text[k] == ' ' {
		k++
	}

	if k > i {
		result = ""
	} else {
		result = text[k : i+1]
	}

	return
}
