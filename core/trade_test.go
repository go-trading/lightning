package core

import (
	"reflect"
	"testing"
	"time"
)

func TestTrade_GetUTCTime(t *testing.T) {
	tests := []struct {
		name string
		Time uint64
		want time.Time
	}{{
		Time: 1585457915462000001,
		want: time.Date(2020, 3, 29, 4, 58, 35, 462000001, time.UTC),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trade := &Trade{Time: tt.Time}
			if got := trade.GetUTCTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trade.GetUTCTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
