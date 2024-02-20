package tyche

import "time"

type Card struct {
	Number [16]int
	Cvv    [3]int
	Exp    time.Time
}
