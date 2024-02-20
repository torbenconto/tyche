package tyche

import (
	"testing"
	"time"
)

func TestCard_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		number int
		exp    string
		want   bool
	}{
		{
			name:   "valid card",
			number: 79927398713,
			exp:    "2021-12-01",
			want:   true,
		},
		{
			name:   "invalid card",
			number: 79927398712,
			exp:    "2021-12-01",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp, _ := time.Parse("2006-01-02", tt.exp)
			c := NewCard(tt.number, exp)
			if got := c.IsValid(); got != tt.want {
				t.Errorf("Card.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
