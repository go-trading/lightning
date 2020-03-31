package bitfinex

import (
	"context"
	"errors"

	"github.com/go-trading/lightning/core"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("p", "bitfinex")
)

func Register(availableServices core.AvailableServices) {
	availableServices.Register("bitfinex", &Bitfinex{})
}

type Bitfinex struct {
	node    *core.Node
	config  *core.ServiceConfig
	symbols core.SymbolsCache
}

func (c *Bitfinex) Name() string {
	return "bitfinex"
}

func (b *Bitfinex) Init(node *core.Node, config *core.ServiceConfig) error {
	log.Trace("Bitfinex.Init")
	b.config = config
	if b.config == nil {
		log.Error("The config is not difined")
		return errors.New("The config is not difined")
	}

	b.node = node
	if node.Type() != core.ProdNode {
		log.Error("Bitfinex is a real exchange")
		return errors.New("Bitfinex is a real exchange")
	}
	return nil
}

func (c *Bitfinex) Start() error {
	return nil
}

func (c *Bitfinex) Stop() error {
	return nil
}

func (c *Bitfinex) Status() core.ServiceStatus {
	return core.Stopped
}

func (b *Bitfinex) SubscribeTrades(ctx context.Context, tt *core.Trades) {}
func (b *Bitfinex) UnsubscribeTrades(tt *core.Trades)                    {}
