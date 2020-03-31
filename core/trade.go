package core

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type (
	ITrade interface {
		GetID() uint64
		GetUTCTime() time.Time
		GetUnixNano() uint64
		GetPrice() MultipliedBy1e8
		GetQTY() MultipliedBy1e8
		//work with csv
		GetCsvHead() string
		GetCsv() string
		InitFromCSV([]string) (howManyLinesDidIRead int, err error)
	}

	Trade struct {
		ID    uint64
		Time  uint64 //UnixNano
		Price MultipliedBy1e8
		QTY   MultipliedBy1e8
	}
)

func (t *Trade) Init(id uint64, time uint64, price MultipliedBy1e8, qty MultipliedBy1e8) {
	t.ID = id
	t.Time = time
	t.Price = price
	t.QTY = qty
}

func (t *Trade) GetUTCTime() time.Time {
	return time.Unix(0, int64(t.Time)).In(time.UTC)
}

func (t *Trade) GetUnixNano() uint64 {
	return t.Time
}

func (t *Trade) GetID() uint64 {
	return t.ID
}

func (t *Trade) GetPrice() MultipliedBy1e8 {
	return t.Price
}

func (t *Trade) GetQTY() MultipliedBy1e8 {
	return t.QTY
}

func (t *Trade) GetCsvHead() string {
	return "id,time,price,qty" //TODO add link for docs
}

func (t *Trade) GetCsv() string {
	return fmt.Sprintf("%v,%v,%v,%v",
		t.ID,
		t.Time,
		t.Price,
		t.QTY,
	)
}

func (t *Trade) InitFromCSV(input []string) (int, error) {
	if len(input) < 4 {
		log.WithField("input fields count", len(input)).Error("can't parse csv with < 4 field")
		return 0, errors.New("can't parse csv with < 4 field")
	}
	id, err := strconv.ParseInt(input[0], 10, 64)
	if err != nil {
		log.WithError(err).WithField("input", input[0]).Error("id parse error")
		return 0, err
	}
	time, err := strconv.ParseInt(input[1], 10, 64)
	if err != nil {
		log.WithError(err).WithField("input", input[1]).Error("time parse error")
		return 0, err
	}
	price, err := strconv.ParseInt(input[2], 10, 64)
	if err != nil {
		log.WithError(err).WithField("input", input[2]).Error("price parse error")
		return 0, err
	}
	qty, err := strconv.ParseInt(input[3], 10, 64)
	if err != nil {
		log.WithError(err).WithField("input", input[3]).Error("qty parse error")
		return 0, err
	}

	t.ID = uint64(id)
	t.Time = uint64(time)
	t.Price = MultipliedBy1e8(price)
	t.QTY = MultipliedBy1e8(qty)
	return 4 /*howManyLinesDidIRead*/, nil
}
