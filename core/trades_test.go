package core

import (
	"testing"
)

func TestTrades_Add_CutTrades(t *testing.T) {
	tt := &Trades{}
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
