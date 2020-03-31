package core

import (
	"context"
	"sync"
	"time"
)

type (
	Trades struct {
		sync.RWMutex
		symbol             *Symbol
		trades             []ITrade
		subscribers        []TradeSubscriber
		cancelSubscription context.CancelFunc
	}

	TradeSubscriber func(*Trades, ITrade)
)

func NewTrades(symbol *Symbol) *Trades {
	return &Trades{
		symbol: symbol,
	}
}

func (tt *Trades) Symbol() *Symbol {
	return tt.symbol
}

func (tt *Trades) SetSymbol(s *Symbol) error {
	if tt.symbol != nil {
		return logAndReturnError("the symbol field can only be set once")
	}
	tt.symbol = s
	return nil
}

func (tt *Trades) Add(trade ITrade) {
	timeStart := time.Now()
	defer func() { addTradesSeconds.Add(float64(time.Now().Sub(timeStart)) / 1e9) }()

	tt.Lock()
	defer tt.Unlock()

	tt.trades = append(tt.trades, trade)

	for _, fn := range tt.subscribers {
		if fn != nil {
			fn(tt, trade)
		}
	}
}

func (tt *Trades) Subscribe(fn TradeSubscriber) int {
	log.WithField("symbol", tt.Symbol().Name()).Debugf("subscribe to trades")
	tt.Lock()
	defer tt.Unlock()

	defer func() {
		if tt.cancelSubscription == nil && sliceContainsNotNil(tt.subscribers) {
			var ctx context.Context
			ctx, tt.cancelSubscription = context.WithCancel(context.TODO())
			tt.Symbol().Exchange().SubscribeTrades(ctx, tt)
		}
	}()

	//try insert in place of removed function
	for i := 0; i < len(tt.subscribers); i++ {
		if tt.subscribers[i] == nil {
			tt.subscribers[i] = fn
			return i
		}
	}
	// insert fn in new place
	tt.subscribers = append(tt.subscribers, fn)
	return len(tt.subscribers) - 1
}

func (tt *Trades) Unsubscribe(id int) {
	log.WithField("id", id).Debug("unsubscribe trades")
	tt.Lock()
	defer tt.Unlock()

	tt.subscribers[id] = nil
	subscriberExists := sliceContainsNotNil(tt.subscribers)
	log.WithField("subscriberExists", subscriberExists).Trace("Unsubscribe?")
	if tt.cancelSubscription != nil && !subscriberExists {
		log.Debug("send cancel to exchange")
		tt.cancelSubscription()
		tt.cancelSubscription = nil
	}
}
