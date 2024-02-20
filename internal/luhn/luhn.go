package luhn

func LuhnCheck(number int) bool {
	for i := number - 1; i >= 0; i-- {
		if i%2 == 0 {
			continue
		}
		if i*2 > 9 {
			i = i*2 - 9
		} else {
			i = i * 2
		}
	}

	return number%10 == 0
}
