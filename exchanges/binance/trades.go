package binance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-trading/lightning/core"
)

func (b *Binance) SubscribeTrades(ctx context.Context, tt *core.Trades) {
	log.WithField("symbol", tt.Symbol().Name()).Debug("subscribe to trades")
	b.waitStop.Add(1)
	go func() {
		defer b.waitStop.Done()
		for {
			err := connectWs(
				ctx,
				fmt.Sprintf("%s@trade", strings.ToLower(tt.Symbol().Name())),
				b.processMessage,
			)
			if errors.Is(err, context.Canceled) {
				log.Info("SubscribeTrades is canceled")
				return
			}
			restartWebsocketCount.Inc()
			if err != nil {
				log.WithError(err).Error("ws error")
			}
		}
	}()
}

func (b *Binance) processMessage(msg []byte) {
	receiveTime := time.Now()
	t := &Trade{}
	err := json.Unmarshal(msg, t)
	if err != nil {
		log.WithError(err).WithField("msg", string(msg)).Trace("error unmarshal trade")
		return
	}
	t.ReceiveTime = receiveTime.UnixNano()
	b.Symbol(t.Symbol).Trades().Add(t)
	messageProcessingTime.Add(float64(time.Now().Sub(receiveTime)) / 1e9)
	messageProcessingCount.Inc()
}
