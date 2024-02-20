package tyche

import (
	"time"
	"tyche/internal/luhn"
)

type Tyche interface {
	IsValid() bool
}

type Card struct {
	Number int
	Cvv    int
	Exp    time.Time
}

func NewCard(number int, cvv int, exp time.Time) *Card {
	return &Card{
		Number: number,
		Cvv:    cvv,
		Exp:    exp,
	}
}

func (c *Card) IsValid() bool {
	return luhn.LuhnCheck(c.Number)
}
