package tyche

func identifyProvider(number string) Provider {
	if len(number) == 0 {
		return Unknown
	}

	firstDigit := number[0]
	secondDigit := rune(0) // Placeholder for the second digit if applicable
	if len(number) > 1 {
		secondDigit = rune(number[1])
	}

	switch firstDigit {
	case '4':
		return Visa
	case '5':
		return MasterCard
	case '3':
		if secondDigit == '4' || secondDigit == '7' {
			return AmericanExpress
		} else if secondDigit == '5' {
			return JCB
		}
	case '6':
		if secondDigit >= '0' && secondDigit <= '5' {
			return Discover
		} else if secondDigit >= '4' && secondDigit <= '6' {
			return Maestro
		} else if secondDigit == '8' || secondDigit == '9' {
			return DinersClub
		}
	}

	return Unknown
}
