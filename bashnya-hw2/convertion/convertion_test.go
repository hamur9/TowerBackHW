package convertion

import (
	"bashnya-hw2/constants"
	"errors"
	"testing"
)

func TestNumberConversion(t *testing.T) {
	testTable := []struct {
		numberIn     int
		resultNumber int
	}{
		{numberIn: -12307, resultNumber: 12308},
		{numberIn: -122, resultNumber: 30736},
		{numberIn: 18, resultNumber: 0},
		{numberIn: 303, resultNumber: 24823},
		{numberIn: 12083, resultNumber: 36256},
		{numberIn: -2882, resultNumber: 0},
	}

	for _, testCase := range testTable {
		result, errCode := NumberConversion(testCase.numberIn)

		t.Logf("Calling NumberConvertion(%d), result = %d\n", testCase.numberIn, result)

		if result != testCase.resultNumber && !errors.Is(errCode, constants.ErrorService) {
			t.Errorf("Incorrect number conversion result.\n Expected:\nNumber = %d\n Received:\nNumber = %d\n",
				testCase.resultNumber, result)
		}
	}
}
