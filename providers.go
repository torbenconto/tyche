package tyche

type Provider string

const (
	Visa            Provider = "Visa"
	MasterCard      Provider = "MasterCard"
	AmericanExpress Provider = "AmericanExpress"
	Discover        Provider = "Discover"
	DinersClub      Provider = "DinersClub"
	JCB             Provider = "JCB"
	Maestro         Provider = "Maestro"
	Unknown         Provider = "Unknown"
)
