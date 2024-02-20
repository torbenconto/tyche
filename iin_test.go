package tyche

import (
	"testing"
)

func TestIdentifyProvider(t *testing.T) {
	tests := []struct {
		name   string
		number string
		want   Provider
	}{
		{
			name:   "visa",
			number: "4111111111111111",
			want:   Visa,
		},
		{
			name:   "mastercard",
			number: "5555555555554444",
			want:   MasterCard,
		},
		{
			name:   "american express",
			number: "378282246310005",
			want:   AmericanExpress,
		},
		{
			name:   "discover",
			number: "6011111111111117",
			want:   Discover,
		},
		{
			name:   "diners club",
			number: "30569309025904",
			want:   DinersClub,
		},
		{
			name:   "jcb",
			number: "3530111333300000",
			want:   JCB,
		},
		{
			name:   "unknown",
			number: "1234567890123456",
			want:   Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := identifyProvider(tt.number); got != tt.want {
				t.Errorf("IdentifyProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
