package tyche

import (
	"time"
	"tyche/internal/validator"
)

type Tyche interface {
	IsValid() bool
	Provider() Provider
}

type Card struct {
	Number string
	Cvv    string
	Exp    time.Time
}

func NewCard(number, cvv string, exp time.Time) *Card {
	return &Card{
		Number: number,
		Cvv:    cvv,
		Exp:    exp,
	}
}

func (c *Card) IsValid() bool {
	return luhnCheck(c.Number) && validator.IsExpired(c.Exp) && len(c.Cvv) == 3
}

func (c *Card) Provider() Provider {
	return identifyProvider(c.Number)
}
