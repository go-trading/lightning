package core

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	addTradesSeconds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "core_add_trades_seconds",
		Help: "The total number of seconds for add trades",
	})
)
