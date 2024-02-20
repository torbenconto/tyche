package tyche

import "strings"

func identifyProvider(number string) Provider {
	if len(number) == 0 {
		return Unknown
	}

	switch {
	case matchesAnyPrefix(number, "34", "37"):
		return AmericanExpress
	case matchesAnyPrefix(number, "5610") || isInBetween(number, "560221", "560225"):
		return BankCard
	case matchesAnyPrefix(number, "604400", "627170", "603522", "589657") ||
		isInBetween(number, "604201", "604219") ||
		isInBetween(number, "604300", "604399"):
		return Cabal
	case isInBetween(number, "300", "305") && len(number) == 14,
		matchesAnyPrefix(number, "2014", "2149"),
		(isInBetween(number, "300", "305") || number[:3] == "309") ||
			(matchesAnyPrefix(number, "36", "38", "39") && len(number) <= 14):
		return DinersClub
	case matchesAnyPrefix(number, "6011") || isInBetween(number, "622126", "622925") ||
		isInBetween(number, "644", "649") || number[:2] == "65":
		return Discover
	case matchesAnyPrefix(number, "4011", "4576"),
		matchesAnyPrefix(number, "431274", "438935", "451416", "457393", "457631", "457632", "504175", "627780", "636297", "636368", "636369"),
		isInBetween(number, "506699", "506778"),
		isInBetween(number, "509000", "509999"),
		isInBetween(number, "650031", "650051"),
		isInBetween(number, "650035", "650033"),
		isInBetween(number, "650405", "650439"),
		isInBetween(number, "650485", "650538"),
		isInBetween(number, "650541", "650598"),
		isInBetween(number, "650700", "650718"),
		isInBetween(number, "650720", "650727"),
		isInBetween(number, "650901", "650920"),
		isInBetween(number, "651652", "651679"),
		isInBetween(number, "655000", "655019"),
		isInBetween(number, "655021", "655021"):
		return Elo
	case matchesAnyPrefix(number, "606282", "637095", "637568", "637599", "637609", "637612"):
		return Hipercard
	case number[:3] == "636" && len(number) >= 16 && len(number) <= 19:
		return InterPayment
	case isInBetween(number, "637", "639"):
		return InstaPayment
	case isInBetween(number, "3528", "3589"):
		return JCB
	case number[:6] == "589562":
		return Naranja
	case matchesAnyPrefix(number, "0604", "5018", "5020", "5038", "5612", "5893", "6304", "6759", "6761", "6762", "6763", "6390"):
		return Maestro
	case number[:4] == "5019":
		return Dankort
	case isInBetween(number, "51", "55") || isInBetween(number, "222100", "272099"):
		return MasterCard
	case matchesAnyPrefix(number, "4026", "4405", "4508", "4844", "4913", "4917", "417500"):
		return VisaElectron
	case number[:1] == "4":
		return Visa
	case number[:2] == "50":
		return Aura
	default:
		return Unknown
	}
}

func matchesAnyPrefix(number string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(number, prefix) {
			return true
		}
	}
	return false
}

func isInBetween(number, start, end string) bool {
	if number[:len(start)] >= start && number[:len(end)] <= end {
		return true
	}

	return false
}
