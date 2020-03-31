package binance

import (
	"reflect"
	"testing"

	"github.com/go-trading/lightning/core"
	"github.com/sirupsen/logrus"
)

func TestBinance_Symbol(t *testing.T) {
	b := &Binance{}

	tests := []struct {
		name       string
		symbolName string
		want       *core.Symbol
	}{
		{
			name:       "incorrect name return nil",
			symbolName: "fkmksdfmlsdmfksd",
			want:       nil,
		},
		{
			name:       "incorrect name return nil",
			symbolName: "BTCUSDT",
			want:       core.NewSymbol(b, "BTCUSDT"),
		},
	}

	logrus.SetLevel(logrus.TraceLevel)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := b.Symbol(tt.symbolName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Binance.Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
