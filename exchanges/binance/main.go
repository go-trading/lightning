package binance

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/go-trading/lightning/core"
)

var (
	log = logrus.WithField("p", "binance")
)

func Register(availableServices core.AvailableServices) {
	availableServices.Register("binance", &Binance{})
}

type Binance struct {
	node     *core.Node
	config   *core.ServiceConfig
	symbols  core.SymbolsCache
	waitStop sync.WaitGroup
}

func (b *Binance) Name() string {
	return "binance"
}

func (b *Binance) Init(node *core.Node, config *core.ServiceConfig) error {
	log.Trace("Binance.Init")
	b.config = config
	if b.config == nil {
		log.Error("The config is not difined")
		return errors.New("The config is not difined")
	}

	b.node = node
	if node.Type() != core.ProdNode {
		log.Error("Binance is a real exchange")
		return errors.New("Binance is a real exchange")
	}
	return nil
}

func (b *Binance) Start() error {
	log.Trace("Binance.Start")
	if b.node.Type() != core.ProdNode {
		log.Error("Binance is a real exchange")
		return errors.New("Binance is a real exchange")
	}
	return nil
}

func (b *Binance) Stop() error {
	log.Trace("Binance.Stoping")
	b.waitStop.Wait()
	log.Trace("Binance.Stoped")
	return nil
}

func (b *Binance) Status() core.ServiceStatus {
	return core.Stopped
}
