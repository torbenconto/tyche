package luhn

import (
	"strconv"
)

func LuhnCheck(number string) bool {
	cardNumber, err := strconv.Atoi(number)
	if err != nil {
		return false
	}

	sum := 0
	double := false
	for cardNumber > 0 {
		digit := cardNumber % 10
		cardNumber /= 10
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}
