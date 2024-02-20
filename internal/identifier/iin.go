package identifier

import "tyche"

func IdentifyProvider(number string) tyche.Provider {
	if len(number) == 0 {
		return tyche.Unknown
	}

	firstDigit := number[0]
	secondDigit := rune(0) // Placeholder for the second digit if applicable
	if len(number) > 1 {
		secondDigit = rune(number[1])
	}

	switch firstDigit {
	case '4':
		return tyche.Visa
	case '5':
		return tyche.MasterCard
	case '3':
		if secondDigit == '4' || secondDigit == '7' {
			return tyche.AmericanExpress
		} else if secondDigit == '5' {
			return tyche.JCB
		}
	case '6':
		if secondDigit >= '0' && secondDigit <= '5' {
			return tyche.Discover
		} else if secondDigit >= '4' && secondDigit <= '6' {
			return tyche.Maestro
		} else if secondDigit == '8' || secondDigit == '9' {
			return tyche.DinersClub
		}
	}

	return tyche.Unknown
}
