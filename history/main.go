package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/go-trading/lightning/core"
	"github.com/go-trading/lightning/exchanges/binance"
	"github.com/go-trading/lightning/history/saver"
	"github.com/go-trading/lightning/web"
)

var log = logrus.WithField("p", "main")

//GithubSHA inject
var GithubSHA string

func main() {
	app := &cli.App{
		Name:    "datasaver",
		Usage:   "datasaver ",
		Action:  StartNode,
		Version: GithubSHA,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func StartNode(cliContext *cli.Context) error {
	if err := core.InitLog(cliContext); err != nil {
		log.Panic(err)
	}

	services := core.NewAvailableServices()
	binance.Register(services)
	saver.Register(services)
	web.Register(services)

	config := core.NodeConfig{
		{
			"service": "binance",
			"key":     "xxxx",
		},
		/*{
			"service": "bitfinex",
		},
				{
				"service":             "arbitration",
				"MinutesForAvgSpread": 5,
				"symbol1":             "binance: BTCUSDT",
				"symbol2":             "bitfinex: tBTCUSD",
			},*/
		{
			"service": "saver",
			"symbols": []string{"binance: BTCUSDT"}, //, "binance: BTCUSDT"},
			"path":    "./saver-data",
		},
		{
			"service": "web",
			"port":    8080,
		},
	}

	node, err := core.NewNode(services, config)
	if err != nil {
		log.WithError(err).Error("Can't init node")
		return err
	}
	err = node.Start()
	if err != nil {
		log.WithError(err).Error("Can't start node")
		return err
	}

	stop := make(chan struct{})
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		log.Info("Got interrupt, shutting down...")
		go func() {
			log.Info("Stopping node")
			node.Stop()
			close(stop)
		}()
		for i := 10; i > 0; i-- {
			<-sigc
			if i > 1 {
				log.Info("Already shutting down, interrupt more to panic", "times", i-1)
			}
		}
		panic("Panic closing the node")
	}()

	<-stop
	return nil
}
