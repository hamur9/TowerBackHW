package text

import "testing"

func TestRemoveLeadAndEndSpaces(t *testing.T) {
	testTable := []struct {
		textIn  string
		textOut string
	}{
		{textIn: "    ноль", textOut: "ноль"},
		{textIn: "один     ", textOut: "один"},
		{textIn: "      пять      ", textOut: "пять"},
		{textIn: "десять", textOut: "десять"},
		{textIn: "            ", textOut: ""},
		{textIn: "", textOut: ""},
	}

	for _, testCase := range testTable {
		resultText := removeLeadAndEndSpaces(testCase.textIn)

		t.Logf("Calling replaceLeadAndEndSpaces, result = %s\n", resultText)

		if resultText != testCase.textOut {
			t.Errorf("len1: %d, len2: %d", len(testCase.textOut), len(resultText))
			t.Errorf("Incorrect replace lead and end spaces result.\nExpected:\nResult = %s\nReceived:\nResult = %s\n",
				testCase.textOut, resultText)
		}
	}
}

func TestNumToText(t *testing.T) {
	testTable := []struct {
		numberIn int
		textOut  string
	}{
		{numberIn: 0, textOut: "ноль"},
		{numberIn: 1, textOut: "один"},
		{numberIn: 5, textOut: "пять"},
		{numberIn: 10, textOut: "десять"},
		{numberIn: 11, textOut: "одиннадцать"},
		{numberIn: 15, textOut: "пятнадцать"},
		{numberIn: 20, textOut: "двадцать"},
		{numberIn: 25, textOut: "двадцать пять"},
		{numberIn: 100, textOut: "сто"},
		{numberIn: 123, textOut: "сто двадцать три"},
		{numberIn: 200, textOut: "двести"},
		{numberIn: 345, textOut: "триста сорок пять"},
		{numberIn: 1000, textOut: "одна тысяча"},
		{numberIn: 1001, textOut: "одна тысяча один"},
		{numberIn: 1015, textOut: "одна тысяча пятнадцать"},
		{numberIn: 1111, textOut: "одна тысяча сто одиннадцать"},
		{numberIn: 2000, textOut: "две тысячи"},
		{numberIn: 2345, textOut: "две тысячи триста сорок пять"},
		{numberIn: 10000, textOut: "десять тысяч"},
		{numberIn: 12345, textOut: "двенадцать тысяч триста сорок пять"},
		{numberIn: 123456, textOut: "сто двадцать три тысячи четыреста пятьдесят шесть"},
	}

	for _, testCase := range testTable {
		resultText := NumToText(testCase.numberIn)

		t.Logf("Calling NumToText(), result = %s\n", resultText)

		if resultText != testCase.textOut {
			t.Errorf("Incorrect num to textOut result.\nExpected:\nResult = %s\nLen = %d\nReceived:\nResult = %s\nLen = %d\n",
				testCase.textOut, len(testCase.textOut), resultText, len(resultText))
		}
	}
}
