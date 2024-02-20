package luhn

func LuhnCheck(number int) bool {
	sum := 0
	double := false
	for number > 0 {
		digit := number % 10
		number /= 10
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
