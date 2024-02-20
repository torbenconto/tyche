package tyche

type Provider string

const (
	Visa            Provider = "Visa"
	MasterCard      Provider = "MasterCard"
	AmericanExpress Provider = "AmericanExpress"
	Discover        Provider = "Discover"
	DinersClub      Provider = "DinersClub"
	JCB             Provider = "JCB"
	Unknown         Provider = "Unknown"
)
