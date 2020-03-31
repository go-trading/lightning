package arbitration

import (
	"errors"

	"github.com/go-trading/lightning/core"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	node                  *core.Node
	MinutesForAvgSpread   uint
	symbol1Name           string
	symbol2Name           string
	trades1               *core.Trades
	trades2               *core.Trades
	trades1SubscriptionId int
	trades2SubscriptionId int
}

var (
	log = logrus.WithField("p", "arbitration")
)

func Register(availableServices core.AvailableServices) {
	availableServices.Register("arbitration", &Bot{})
}

func (b *Bot) Name() string {
	return "Arbitration bot"
}

func (b *Bot) Init(node *core.Node, config *core.ServiceConfig) error {
	log.Debug("Bot.Init")
	b.node = node
	if b.node == nil {
		log.Error("The node is not difined")
		return errors.New("The node is not difined")
	}

	b.MinutesForAvgSpread = config.GetUint("MinutesForAvgSpread")
	b.symbol1Name = config.GetString("symbol1")
	b.symbol2Name = config.GetString("symbol2")
	if err := config.GetErrors(); err != nil {
		log.WithError(err).Error("maybe incorrect arbitration config")
		return err
	}
	return nil
}

func (b *Bot) Start() error {
	symbol1 := b.node.Symbol(b.symbol1Name)
	symbol2 := b.node.Symbol(b.symbol2Name)

	if symbol1 == nil || symbol2 == nil {
		log.WithFields(logrus.Fields{
			"symbol1": symbol1,
			"symbol2": symbol2,
		}).Error("symbols not found")
		return errors.New("symbols not found")
	}
	if symbol1.Exchange() == symbol2.Exchange() && symbol1.Name() == symbol2.Name() {
		log.Error("Arbitration beetween equal symbols")
		return errors.New("Arbitration beetween equal symbols")
	}

	b.trades1 = symbol1.Trades()
	b.trades2 = symbol2.Trades()

	b.trades1SubscriptionId = b.trades1.Subscribe(b.onNewTrade)
	b.trades2SubscriptionId = b.trades2.Subscribe(b.onNewTrade)

	return nil
}

func (b *Bot) Stop() error {
	b.trades1.Unsubscribe(b.trades1SubscriptionId)
	b.trades2.Unsubscribe(b.trades2SubscriptionId)
	return nil
}

func (b *Bot) Status() core.ServiceStatus {
	return core.Stopped
}

func (b *Bot) SubscribeStatus(func(core.ServiceStatus))   {}
func (b *Bot) UnsubscribeStatus(func(core.ServiceStatus)) {}
