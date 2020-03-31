package core

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	data := []byte(`
  - service: binance
    key: xxxx
  - service: coinbase
  - service: arbitration
    minutes_for_avg: 5
  - service: web
    port: 8080
`)

	want := &NodeConfig{
		{
			"service": "binance",
			"key":     "xxxx",
		},
		{
			"service": "coinbase",
		},
		{
			"service":         "arbitration",
			"minutes_for_avg": 5,
		},
		{
			"service": "web",
			"port":    8080,
		},
	}

	if got := ParseConfig(data); !reflect.DeepEqual(got, want) {
		t.Errorf("ParseConfig() = %v, want %v", got, want)
	}
}
