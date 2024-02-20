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
			name:   "maestro",
			number: "6759649826438453",
			want:   Maestro,
		},
		{
			name:   "visa electron",
			number: "4917300800000000",
			want:   VisaElectron,
		},
		{
			name:   "aura",
			number: "5078601870000127985",
			want:   Aura,
		},
		{
			name:   "cabal",
			number: "5896570000000000",
			want:   Cabal,
		},
		{
			name:   "elo",
			number: "5066991111111118",
			want:   Elo,
		},
		{

			name:   "hipercard",
			number: "6062825624254001",
			want:   Hipercard,
		},
		{
			name:   "interpayment",
			number: "6360111331111111",
			want:   InterPayment,
		},

		{
			name:   "instapayment",
			number: "6388000000000000",

			want: InstaPayment,
		},
		{
			name:   "naranja",
			number: "5895627820000000",
			want:   Naranja,
		},
		{
			name:   "dankort",
			number: "5019717010103742",
			want:   Dankort,
		},
		{
			name:   "bank card",
			number: "5610591081018250",
			want:   BankCard,
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
