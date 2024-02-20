package luhn

import "testing"

func TestLuhnCheck(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "valid number",
			number: 79927398713,
			want:   true,
		},
		{
			name:   "invalid number",
			number: 79927398712,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LuhnCheck(tt.number); got != tt.want {
				t.Errorf("LuhnCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
