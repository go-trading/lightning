package binance

import (
	"encoding/json"
)

func (b *Binance) clientGetExchangeInfo() (*ExchangeInfo, error) {
	body, err := b.clientGet("/api/v3/exchangeInfo")
	if err != nil {
		log.WithError(err).Error("can't get ExchangeInfo")
		return nil, err
	}

	exchangeInfo := &ExchangeInfo{}
	err = json.Unmarshal(body, &exchangeInfo)
	if err != nil {
		log.WithError(err).WithField("body", string(body)).Error("can't unmarshal ExchangeInfo")
		return nil, err
	}

	return exchangeInfo, nil
}

type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []interface{} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []Symbols     `json:"symbols"`
}

type Symbols struct {
	Symbol                 string        `json:"symbol"`
	Status                 string        `json:"status"`
	BaseAsset              string        `json:"baseAsset"`
	BaseAssetPrecision     int           `json:"baseAssetPrecision"`
	QuoteAsset             string        `json:"quoteAsset"`
	QuotePrecision         int           `json:"quotePrecision"`
	OrderTypes             bool          `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed bool          `json:"isMarginTradingAllowed"`
	Filters                []interface{} `json:"filters"`
}
