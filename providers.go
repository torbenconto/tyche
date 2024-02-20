package tyche

type Provider string

const (
	Aura            Provider = "Aura"
	AmericanExpress Provider = "AmericanExpress"
	BankCard        Provider = "BankCard"
	Cabal           Provider = "Cabal"
	Dankort         Provider = "Dankort"
	DinersClub      Provider = "DinersClub"
	Discover        Provider = "Discover"
	Elo             Provider = "Elo"
	Hipercard       Provider = "Hipercard"
	InstaPayment    Provider = "InstaPayment"
	InterPayment    Provider = "InterPayment"
	JCB             Provider = "JCB"
	Maestro         Provider = "Maestro"
	MasterCard      Provider = "MasterCard"
	Naranja         Provider = "Naranja"
	Unknown         Provider = "Unknown"
	Visa            Provider = "Visa"
	VisaElectron    Provider = "VisaElectron"
)

func (p Provider) String() string {
	return string(p)
}
