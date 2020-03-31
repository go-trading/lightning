package bitfinex

import (
	"encoding/json"

	"github.com/go-trading/lightning/core"
)

func (b *Bitfinex) clientGetTickers() (map[string]*core.Symbol, error) {
	body, err := b.clientGetPub("/v2/tickers?symbols=ALL")
	if err != nil {
		log.WithError(err).Error("can't get tickers")
		return nil, err
	}
	tickers := make([][]interface{}, 0, 500)
	err = json.Unmarshal(body, &tickers)
	if err != nil {
		log.WithError(err).WithField("body", string(body)).Error("can't unmarshal ExchangeInfo")
		return nil, err
	}

	symbols := make(map[string]*core.Symbol)
	for _, ticker := range tickers {
		sName, ok := ticker[0].(string)
		if !ok {
			log.WithField("ticker", ticker).Error("Can't read ticker name")
		} else {
			symbols[sName] = core.NewSymbol(b, sName)
		}
	}
	return symbols, nil
}
