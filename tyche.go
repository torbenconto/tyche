package tyche

import (
	"time"
	"tyche/internal/validator"
)

type Tyche interface {
	IsValid() bool
	IdentifyProvider() Provider
}

type Card struct {
	Number string
	Exp    time.Time
}

func NewCard(number string, exp time.Time) *Card {
	return &Card{
		Number: number,
		Exp:    exp,
	}
}

func (c *Card) IsValid() bool {
	return luhnCheck(c.Number) && validator.IsExpired(c.Exp)
}

func (c *Card) IdentifyProvider() Provider {
	return identifyProvider(c.Number)
}
