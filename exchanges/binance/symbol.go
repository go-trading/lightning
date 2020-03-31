package binance

import (
	"github.com/go-trading/lightning/core"
)

func (b *Binance) Symbol(symbolName string) *core.Symbol {
	if b.symbols.IsCacheInvalid(b.node.Now()) {
		exchangeInfo, err := b.clientGetExchangeInfo()
		if err != nil {
			log.WithError(err).Error("Can't clientGetExchangeInfo")
		} else {
			newSymbols := make(map[string]*core.Symbol)
			for _, s := range exchangeInfo.Symbols {
				newSymbols[s.Symbol] = core.NewSymbol(b, s.Symbol)
			}
			b.symbols.Add(newSymbols, b.node.Now())
		}
	}
	return b.symbols.Get(symbolName)
}
