package saver

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	queueLength = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "saver_trades_queue_length",
		Help: "The current number of trades in queue waiting to be writed",
	})
)
