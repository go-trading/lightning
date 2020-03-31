package bitfinex

import (
	"github.com/go-trading/lightning/core"
)

func (b *Bitfinex) Symbol(symbolName string) *core.Symbol {
	if b.symbols.IsCacheInvalid(b.node.Now()) {
		tickers, err := b.clientGetTickers()
		if err != nil {
			log.WithError(err).Error("Can't clientGetExchangeInfo")
		} else {
			b.symbols.Add(tickers, b.node.Now())
		}
	}
	return b.symbols.Get(symbolName)
}
