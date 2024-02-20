package tyche

import (
	"time"
	"tyche/internal/luhn"
	"tyche/internal/validator"
)

type Tyche interface {
	IsValid() bool
}

type Card struct {
	Number int
	Exp    time.Time
}

func NewCard(number int, exp time.Time) *Card {
	return &Card{
		Number: number,
		Exp:    exp,
	}
}

func (c *Card) IsValid() bool {
	return luhn.LuhnCheck(c.Number) && validator.IsExpired(c.Exp)
}
