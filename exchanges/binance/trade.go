package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-trading/lightning/core"
)

type Trade struct {
	core.Trade
	Symbol        string `json:"s"`
	BuyerOrderID  uint64 `json:"b"`
	SellerOrderID uint64 `json:"a"`
	IsMaker       bool   `json:"m"`
	ReceiveTime   int64
}

func (t *Trade) UnmarshalJSON(data []byte) error {
	type TradeAlias Trade
	aux := &struct {
		TradeTime uint64               `json:"T"`
		Type      string               `json:"e"`
		EventTime int64                `json:"E"`
		TradeID   uint64               `json:"t"`
		Price     core.MultipliedBy1e8 `json:"p"`
		Quantity  core.MultipliedBy1e8 `json:"q"`
		*TradeAlias
	}{
		TradeAlias: (*TradeAlias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Type != "trade" {
		log.WithField("data", string(data)).Error("input data type is not trade")
		return errors.New("input data type is not trade")
	}
	t.Trade.Init(aux.TradeID, aux.TradeTime*1e6, aux.Price, aux.Quantity)

	return nil
}

func (t *Trade) GetCsvHead() string {
	return t.Trade.GetCsvHead() + ",BuyerOrderID,SellerOrderID,IsMaker,ReceiveTime"
}

func (t *Trade) GetCsv() string {
	return fmt.Sprintf("%s,%v,%v,%v,%v",
		t.Trade.GetCsv(),
		t.BuyerOrderID,
		t.SellerOrderID,
		t.IsMaker,
		t.ReceiveTime,
	)
}

func (t *Trade) InitFromCSV(input []string) (int, error) {
	howManyLinesReaded, err := t.Trade.InitFromCSV(input)

	if len(input) < (3 + howManyLinesReaded) {
		log.WithField("input fields count", len(input)).Error("can't parse csv (not enough fields)")
		return 0, errors.New("can't parse csv (not enough fields)")
	}
	buyerOrderID, err := strconv.ParseInt(input[howManyLinesReaded+0], 10, 64)
	if err != nil {
		log.WithError(err).WithField("input", input[howManyLinesReaded+0]).Error("buyerOrderID parse error")
		return howManyLinesReaded, err
	}
	sellerOrderID, err := strconv.ParseInt(input[howManyLinesReaded+1], 10, 64)
	if err != nil {
		log.WithError(err).WithField("input", input[howManyLinesReaded+1]).Error("sellerOrderID parse error")
		return howManyLinesReaded, err
	}
	isMaker, err := strconv.ParseBool(input[howManyLinesReaded+2])
	if err != nil {
		log.WithError(err).WithField("input", input[howManyLinesReaded+2]).Error("isMaker parse error")
		return 0, err
	}

	t.BuyerOrderID = uint64(buyerOrderID)
	t.SellerOrderID = uint64(sellerOrderID)
	t.IsMaker = isMaker
	return howManyLinesReaded + 3, nil
}
