package convertion

import (
	"bashnya-hw2/constants"
)

func NumberConversion(num int) (result int, errorCode error) {
	result = num
	for result < constants.MaxNumber {
		if result < 0 {
			result *= -1
		} else if result%7 == 0 {
			result *= 39
		} else if result%9 == 0 {
			result = result*13 + 1
			continue
		} else {
			result = (result + 2) * 3
		}

		if result%13 == 0 && result%9 == 0 {
			result = 0
			errorCode = constants.ErrorService
			break
		} else {
			result += 1
		}
	}

	return
}
