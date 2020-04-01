package core

import (
	"testing"
)

func TestTrades_Add_CutTrades(t *testing.T) {
	tt := &Trades{}
	tt.INeedHistoricalTrades(1000)
	for i := 0; i < 2500; i++ {
		tt.Add(&Trade{ID: uint64(i)})
	}
	if len(tt.trades) != 1500 {
		t.Errorf("wait len=1500, but have %v", len(tt.trades))
	}
	for i := 0; i < 1500; i++ {
		if tt.trades[i].GetID() != uint64(i)+1000 {
			t.Fatalf("incorect item. wait %v, have %v", i+1000, tt.trades[i].GetID())
		}
	}
}

func TestTrades_INeedHistoricalTrades(t *testing.T) {
	tt := &Trades{}

	tt.Add(&Trade{ID: 1})
	if len(tt.trades) != 0 {
		t.Errorf("don't want save trades, but save")
	}

	tt.INeedHistoricalTrades(2)
	tt.Add(&Trade{ID: 2})
	if len(tt.trades) != 1 {
		t.Errorf("want save 1 trades, but have %v", len(tt.trades))
	}

	tt.INeedHistoricalTrades(1)
	if tt.historicalTradesCount != 2 {
		t.Errorf("reduce historicalTradesCount %v", tt.historicalTradesCount)
	}
}
