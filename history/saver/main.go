package saver

import (
	"os"
	"path"
	"sync"
	"time"

	"github.com/go-trading/lightning/core"
	"github.com/sirupsen/logrus"
)

type (
	Saver struct {
		sync.WaitGroup
		symbols map[string]symbolItem
		path    string
		input   chan inputItem
	}
	symbolItem struct {
		symbol       *core.Symbol
		subId        int
		file         *os.File
		openFileDate time.Time
	}
	inputItem struct {
		tt *core.Trades
		t  core.ITrade
	}
)

var (
	log = logrus.WithField("p", "saver")
)

func Register(availableServices core.AvailableServices) {
	availableServices.Register("saver", &Saver{})
}

func (s *Saver) Name() string {
	return "saver"
}

func (s *Saver) Init(node *core.Node, config *core.ServiceConfig) error {
	log.Debug("saver.Init")
	s.symbols = make(map[string]symbolItem)
	for _, symbolName := range config.GetStrings("symbols") {
		symbol := node.Symbol(symbolName)
		if symbol == nil {
			log.Errorf("symbol %s not found\n", symbolName)
		} else {
			s.symbols[buildKey(symbol)] = symbolItem{symbol: symbol}
		}
	}
	s.path = config.GetString("path")
	if err := config.GetErrors(); err != nil {
		log.WithError(err).Error("maybe incorrect saver config")
		return err
	}
	return nil
}

func (s *Saver) Start() error {
	log.Debug("saver starting")
	s.input = make(chan inputItem, 1000)
	//subscribe to all symbols
	for symbol, _ := range s.symbols {
		saverSymbol := s.symbols[symbol]
		saverSymbol.subId = saverSymbol.symbol.Trades().Subscribe(s.onNewTrade)
	}
	s.write()
	log.WithField("symbols", s.symbols).Info("saver started")
	return nil
}

func (s *Saver) Stop() error {
	log.Debug("saver stoping")
	for _, saverSymbol := range s.symbols {
		saverSymbol.symbol.Trades().Unsubscribe(saverSymbol.subId)
	}
	close(s.input)
	s.Wait()

	for _, saverSymbol := range s.symbols {
		if saverSymbol.file != nil {
			if err := saverSymbol.file.Close(); err != nil {
				log.WithError(err).Error("close file error")
			}
		}
	}

	log.Info("saver stoped")
	return nil
}

func (s *Saver) Status() core.ServiceStatus {
	return core.Stopped
}

func (s *Saver) SubscribeStatus(func(core.ServiceStatus))   {}
func (s *Saver) UnsubscribeStatus(func(core.ServiceStatus)) {}

func (s *Saver) onNewTrade(tt *core.Trades, t core.ITrade) {
	s.input <- inputItem{tt: tt, t: t}
}

func (s *Saver) write() {
	s.Add(1)
	defer s.Done()
	go func() {
		for {
			i, ok := <-s.input
			queueLength.Set(float64(len(s.input)))
			if !ok {
				return
			}
			saverSymbol, ok := s.symbols[buildKey(i.tt.Symbol())]
			if !ok {
				log.WithField("symbol", buildKey(i.tt.Symbol())).Error("symbol not found")
				continue
			}

			y, m, d := i.t.GetUTCTime().Date()
			tradeDate := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
			if saverSymbol.openFileDate != tradeDate || saverSymbol.file == nil {
				//close previos file
				if saverSymbol.file != nil {
					if err := saverSymbol.file.Close(); err != nil {
						log.WithError(err).Error("Close file error")
						continue
					}
				}
				//check path and create if not exists
				filePath := path.Join(
					s.path,
					saverSymbol.symbol.Exchange().Name(),
					saverSymbol.symbol.Name(),
				)
				if _, notExistsError := os.Stat(filePath); os.IsNotExist(notExistsError) {
					if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
						log.WithError(err).Error("can't make directory")
						continue
					}
				}

				//open new date file
				fileName := path.Join(filePath, tradeDate.Format("20060102.csv"))

				needHeader := false
				if _, notExistsError := os.Stat(fileName); os.IsNotExist(notExistsError) {
					needHeader = true
				}

				file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					log.WithError(err).Error("can't open file")
					continue
				}
				saverSymbol.openFileDate = tradeDate
				saverSymbol.file = file
				s.symbols[buildKey(i.tt.Symbol())] = saverSymbol

				if needHeader {
					saverSymbol.file.WriteString(i.t.GetCsvHead() + "\n")
				}
			}

			if _, err := saverSymbol.file.WriteString(i.t.GetCsv() + "\n"); err != nil {
				log.WithError(err).Error("Write")
				continue
			}
		}
	}()
}

func buildKey(symbol *core.Symbol) string {
	return symbol.Exchange().Name() + "-" + symbol.Name()
}
