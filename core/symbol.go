package core

import (
	"sync"
)

type SymbolStatus int32

const (
	SymbolStatusIndefined SymbolStatus = iota
	SymbolStatusStopped
)

type Symbol struct {
	exchange Exchange
	name     string
	trades   *Trades
	// see symbol_status.go
	status                 SymbolStatus
	statusSubscribersMutex sync.RWMutex
	statusSubscribers      []func(*Symbol, SymbolStatus)
}

func NewSymbol(e Exchange, n string) *Symbol {
	s := &Symbol{
		exchange: e,
		name:     n,
	}
	s.trades = NewTrades(s)
	return s
}

func (s *Symbol) Exchange() Exchange {
	return s.exchange
}

func (s *Symbol) Name() string {
	return s.name
}

func (s *Symbol) Trades() *Trades {
	return s.trades
}
