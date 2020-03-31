package binance

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	restartWebsocketCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "binance_restart_websocket_count",
		Help: "The total number of restart websocket",
	})
	websocketWaitNewMessageTime = promauto.NewCounter(prometheus.CounterOpts{
		Name: "binance_ws_wait_new_message_seconds",
		Help: "The total number of second wait new message websocket",
	})
	websocketMessageProcessingTime = promauto.NewCounter(prometheus.CounterOpts{
		Name: "binance_ws_message_processing_seconds",
		Help: "The total number of second spend to processing message from websocket",
	})
	messageProcessingTime = promauto.NewCounter(prometheus.CounterOpts{
		Name: "binance_message_processing_seconds",
		Help: "The total number of second spend to processing message",
	})
	messageProcessingCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "binance_message_count",
		Help: "The total number of processing messages",
	})
)
