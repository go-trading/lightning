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
			"symbols": allBinance,// []string{"binance: BTCUSDT"}, //, "binance: BTCUSDT"},
			"path":    "./data",
		},
		{
			"service": "web",
			"address": ":80",
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

allBinance := []string{"binance: XLMPAX","binance: TRXTUSD","binance: ZRXUSDT","binance: WTCETH","binance: BQXETH","binance: FUELETH","binance: VIABNB","binance: STORMETH","binance: XTZBNB","binance: TRXBTC","binance: NULSBTC","binance: CNDBTC","binance: POLYBTC","binance: GTOPAX","binance: EOSBEARBUSD","binance: COTIBTC","binance: ADXBNB","binance: PPTETH","binance: LINKPAX","binance: ZILUSDT","binance: NEOUSDC","binance: EOSBULLUSDT","binance: ONTBUSD","binance: OAXETH","binance: WTCBTC","binance: CELRUSDT","binance: ATOMUSDT","binance: NPXSUSDT","binance: SCBNB","binance: HOTETH","binance: PAXETH","binance: NULSBNB","binance: DLTBTC","binance: INSBTC","binance: XRPUSDT","binance: ENJBNB","binance: RVNBUSD","binance: NANOBNB","binance: TRXUSDT","binance: DUSKPAX","binance: LRCETH","binance: DASHBTC","binance: CDTETH","binance: GXSBTC","binance: NAVBNB","binance: ANKRUSDC","binance: BEAMBTC","binance: XTZUSDT","binance: TRXBUSD","binance: ICXBUSD","binance: CLOAKBTC","binance: OGNUSDT","binance: LSKUSDT","binance: CNDBNB","binance: ALGOBNB","binance: PERLUSDC","binance: ETHBUSD","binance: LTOBTC","binance: DASHUSDT","binance: QTUMBUSD","binance: ETHBULLUSDT","binance: XLMETH","binance: NANOBTC","binance: BCNETH","binance: IOTXBTC","binance: BTTBNB","binance: REPBTC","binance: AGIETH","binance: NPXSBTC","binance: CELRBTC","binance: EURUSDT","binance: EOSBUSD","binance: ETHBTC","binance: ZRXBTC","binance: EOSBTC","binance: CVCETH","binance: ZECPAX","binance: BTSUSDT","binance: QSPBTC","binance: WAVESETH","binance: TRXXRP","binance: BUSDRUB","binance: VITEBTC","binance: DOGEUSDC","binance: ADXETH","binance: LTCBNB","binance: TUSDBTC","binance: LINKTUSD","binance: USDSUSDC","binance: ATOMBTC","binance: ANKRTUSD","binance: MCOETH","binance: REQETH","binance: NEOUSDT","binance: ADAETH","binance: XEMBNB","binance: CVCBNB","binance: XRPBNB","binance: STXUSDT","binance: SNTBTC","binance: NULSETH","binance: BATBTC","binance: TNBETH","binance: PIVXBNB","binance: BNBBULLBUSD","binance: ADAUSDT","binance: USDCTUSD","binance: DENTUSDT","binance: STXBTC","binance: EVXETH","binance: ARNETH","binance: USDSTUSD","binance: NKNBTC","binance: LTOBNB","binance: VITEBNB","binance: MTLBTC","binance: WAVESUSDT","binance: NPXSUSDC","binance: WINTRX","binance: BEAMUSDT","binance: XEMETH","binance: ZENBNB","binance: HOTBNB","binance: LTCBTC","binance: SALTETH","binance: MDABTC","binance: FUELBTC","binance: ONTBNB","binance: BANDBNB","binance: XRPTRY","binance: STRATETH","binance: SUBBTC","binance: VENBTC","binance: WPRETH","binance: PAXBTC","binance: ICNBTC","binance: PAXUSDT","binance: MITHUSDT","binance: ETHBULLBUSD","binance: BCCBTC","binance: MITHBTC","binance: ADAPAX","binance: COCOSBTC","binance: BNBBEARBUSD","binance: WTCBNB","binance: NXSETH","binance: TRXBNB","binance: KEYUSDT","binance: GTOUSDT","binance: XLMBUSD","binance: MODETH","binance: QTUMBNB","binance: SYSBTC","binance: LINKUSDC","binance: FETBNB","binance: LINKUSDT","binance: MATICBTC","binance: PHBBTC","binance: BTCNGN","binance: CTXCUSDT","binance: ASTBTC","binance: VIABTC","binance: NCASHBNB","binance: THETABTC","binance: TOMOUSDC","binance: ETHUSDT","binance: MCOBTC","binance: XVGBTC","binance: BCHUSDC","binance: ETCETH","binance: ENGETH","binance: WINUSDT","binance: HBARUSDT","binance: XTZBUSD","binance: LINKBTC","binance: VENETH","binance: VENUSDT","binance: BEAMBNB","binance: ONTUSDT","binance: ATOMUSDC","binance: TROYUSDT","binance: NEOBTC","binance: EOSETH","binance: FUNETH","binance: ADXBTC","binance: RLCBTC","binance: BNTBUSD","binance: OSTBTC","binance: ETCBNB","binance: PHBPAX","binance: COSUSDT","binance: NKNBNB","binance: QTUMETH","binance: IOTABNB","binance: ZECUSDC","binance: HBARBTC","binance: VITEUSDT","binance: TUSDBTUSD","binance: LINKBUSD","binance: POWRETH","binance: BCCUSDT","binance: TNTBTC","binance: PHXETH","binance: ENJUSDT","binance: EOSUSDC","binance: SUBETH","binance: CMTBNB","binance: ELFETH","binance: QKCETH","binance: RENBNB","binance: DASHETH","binance: DCRBNB","binance: BCHABCUSDT","binance: XRPBUSD","binance: ICXETH","binance: HOTBTC","binance: EOSTUSD","binance: BCHABCPAX","binance: PHBUSDC","binance: NXSBNB","binance: ATOMTUSD","binance: FTMTUSD","binance: ZECBTC","binance: ZECETH","binance: HSRETH","binance: GTOETH","binance: IOTXETH","binance: NEOBUSD","binance: MBLBNB","binance: BNBBEARUSDT","binance: ICXBNB","binance: ZECUSDT","binance: BATTUSD","binance: GTOUSDC","binance: WANUSDT","binance: ADABTC","binance: CHATBTC","binance: CELRBNB","binance: THETAUSDT","binance: LTCBUSD","binance: BCCETH","binance: BRDBNB","binance: WANBTC","binance: HCBTC","binance: ADAUSDC","binance: CHZBNB","binance: APPCBNB","binance: ONTBTC","binance: IOTAUSDT","binance: COSBNB","binance: TOMOUSDT","binance: POWRBTC","binance: RCNETH","binance: RCNBNB","binance: ARDRBNB","binance: BUSDZAR","binance: BCHUSDT","binance: XRPRUB","binance: AEBNB","binance: WANBNB","binance: BCNBTC","binance: BNBUSDC","binance: ONGUSDT","binance: CVCUSDT","binance: COTIBNB","binance: GNTBTC","binance: XRPPAX","binance: BTCBBTC","binance: BTCUSDT","binance: BCPTBTC","binance: GXSETH","binance: TNTETH","binance: TRIGBTC","binance: BTSBUSD","binance: PHBBNB","binance: BCPTTUSD","binance: CHZUSDT","binance: DNTETH","binance: ENGBTC","binance: LSKBNB","binance: KEYETH","binance: ATOMBNB","binance: TUSDETH","binance: VETETH","binance: KAVAUSDT","binance: SNGLSETH","binance: BNBUSDT","binance: LSKETH","binance: MANAETH","binance: BCNBNB","binance: ADABUSD","binance: LTOUSDT","binance: STPTBTC","binance: BCPTETH","binance: WAVESTUSD","binance: ONEBNB","binance: COSBTC","binance: PERLBNB","binance: TFUELTUSD","binance: ONEUSDC","binance: BTGETH","binance: DLTETH","binance: BTTUSDC","binance: DASHBNB","binance: TFUELBNB","binance: BQXBTC","binance: BCHBUSD","binance: ALGOPAX","binance: COTIUSDT","binance: STRATBTC","binance: SNMETH","binance: BTSBNB","binance: XZCBNB","binance: TRIGBNB","binance: MATICUSDT","binance: ALGOUSDC","binance: STXBNB","binance: ZRXETH","binance: ETCBTC","binance: STEEMETH","binance: GNTBNB","binance: PHXBNB","binance: CTXCBTC","binance: FTTBTC","binance: TCTBNB","binance: ETHZAR","binance: KMDBTC","binance: APPCBTC","binance: KEYBTC","binance: USDCBNB","binance: LTCUSDC","binance: TOMOBTC","binance: ETHTUSD","binance: WAVESPAX","binance: OMGUSDT","binance: MATICBNB","binance: BATUSDC","binance: EVXBTC","binance: KMDETH","binance: LENDBTC","binance: BATBUSD","binance: LTCTUSD","binance: ZRXBNB","binance: NKNUSDT","binance: EURBUSD","binance: AIONBUSD","binance: BCCBNB","binance: ICXUSDT","binance: BCHABCBTC","binance: DUSKUSDT","binance: CTXCBNB","binance: ADABNB","binance: REQBTC","binance: ELFBTC","binance: NANOETH","binance: WANETH","binance: GRSBTC","binance: LINKETH","binance: ERDUSDT","binance: STRATBUSD","binance: RLCETH","binance: MFTETH","binance: OMGBNB","binance: KAVABTC","binance: TOMOBUSD","binance: XRPTUSD","binance: BCHSVUSDC","binance: BANDUSDT","binance: OAXBTC","binance: WAVESBTC","binance: VIAETH","binance: VETBTC","binance: POLYBNB","binance: RVNUSDT","binance: FTTBNB","binance: MTHBTC","binance: QSPBNB","binance: TFUELBTC","binance: DOGEBNB","binance: MFTUSDT","binance: XZCBTC","binance: FTMUSDT","binance: XMRBUSD","binance: NEBLBNB","binance: RENBTC","binance: BCHABCBUSD","binance: USDTRUB","binance: WINUSDC","binance: DREPBTC","binance: TRXETH","binance: STORMBTC","binance: QLCBNB","binance: THETAETH","binance: XMRUSDT","binance: RPXETH","binance: QLCBTC","binance: PHXBTC","binance: ALGOTUSD","binance: NANOBUSD","binance: XMRBNB","binance: ONETUSD","binance: BTTBUSD","binance: XZCETH","binance: ETCUSDT","binance: VETBNB","binance: PAXTUSD","binance: ONGBNB","binance: BUSDUSDT","binance: TCTBTC","binance: ICNETH","binance: AMBBTC","binance: XLMUSDT","binance: PAXBNB","binance: FTMPAX","binance: VENBNB","binance: BTSBTC","binance: AIONBTC","binance: PIVXBTC","binance: ONTUSDC","binance: CHZBTC","binance: ARKBTC","binance: STORJBTC","binance: WABIBTC","binance: LTCETH","binance: REPBNB","binance: QTUMUSDT","binance: DATABTC","binance: BCHSVPAX","binance: BCPTPAX","binance: BNBZAR","binance: TROYBNB","binance: ENJBUSD","binance: PPTBTC","binance: VIBEBTC","binance: ZILBTC","binance: XLMUSDC","binance: ETCUSDC","binance: BATBNB","binance: ONTETH","binance: NPXSETH","binance: STORMUSDT","binance: BULLUSDT","binance: STRATUSDT","binance: NCASHBTC","binance: POAETH","binance: ETCBUSD","binance: ETHRUB","binance: BTCTRY","binance: BCHSVTUSD","binance: KNCBTC","binance: RCNBTC","binance: EDOETH","binance: AGIBTC","binance: NEOTUSD","binance: BCDETH","binance: ANKRPAX","binance: BULLBUSD","binance: XRPBULLBUSD","binance: XRPBEARUSDT","binance: BCDBTC","binance: IOSTETH","binance: RVNBNB","binance: BEARUSDT","binance: BTSETH","binance: DGDETH","binance: OSTETH","binance: MFTBTC","binance: FTMBNB","binance: HOTUSDT","binance: BATUSDT","binance: ENJBTC","binance: MCOBNB","binance: SYSETH","binance: SCETH","binance: BTTTUSD","binance: PERLBTC","binance: BCHBNB","binance: GASBTC","binance: DGDBTC","binance: GTOBTC","binance: LUNBTC","binance: ALGOUSDT","binance: ONEBTC","binance: ATOMBUSD","binance: BATETH","binance: INSETH","binance: LOOMETH","binance: QKCBTC","binance: ARDRETH","binance: YOYOBTC","binance: TFUELPAX","binance: ONTPAX","binance: USDTTRY","binance: ALGOBTC","binance: MTLETH","binance: BRDETH","binance: AGIBNB","binance: XLMTUSD","binance: BNBUSDS","binance: BTCRUB","binance: IOTAETH","binance: VIBBTC","binance: STEEMBTC","binance: HCUSDT","binance: BCHTUSD","binance: KAVABNB","binance: BNBRUB","binance: TCTUSDT","binance: CHATETH","binance: THETABNB","binance: BNBPAX","binance: BCHABCUSDC","binance: BTTUSDT","binance: HCETH","binance: GOBNB","binance: ADATUSD","binance: POWRBNB","binance: RDNETH","binance: POEETH","binance: NASETH","binance: DOCKBTC","binance: HBARBNB","binance: LOOMBTC","binance: REPETH","binance: BTTPAX","binance: BNBETH","binance: DNTBTC","binance:
 GTOBNB","binance: ICXBTC","binance: WPRBTC","binance: DOGEPAX","binance: BGBPUSDC","binance: DOCKUSDT","binance: BATPAX","binance: DUSKUSDC","binance: BCHABCTUSD","binance: ETCTUSD","binance: ANKRBTC","binance: BNTETH","binance: SALTBTC","binance: ASTETH","binance: RDNBTC","binance: LTCUSDT","binance: MTLUSDT","binance: WAVESBUSD","binance: NANOUSDT","binance: DUSKBNB","binance: RLCUSDT","binance: USDTZAR","binance: TFUELUSDC","binance: ERDBTC","binance: BNBNGN","binance: FUNBTC","binance: STORJETH","binance: APPCETH","binance: AEBTC","binance: ATOMPAX","binance: OGNBTC","binance: DREPUSDT","binance: BNBBULLUSDT","binance: SNTETH","binance: ZILETH","binance: CLOAKETH","binance: ETHUSDC","binance: ANKRBNB","binance: VIBETH","binance: WABIBNB","binance: BNBTUSD","binance: ONEUSDT","binance: WINBNB","binance: ETHEUR","binance: DREPBNB","binance: BEARBUSD","binance: OMGBTC","binance: DLTBNB","binance: GRSETH","binance: DENTBTC","binance: DOCKETH","binance: EOSBEARUSDT","binance: XRPBULLUSDT","binance: STRATBNB","binance: ERDPAX","binance: WRXUSDT","binance: BCPTBNB","binance: NEBLETH","binance: AEETH","binance: POABNB","binance: TUSDBNB","binance: KNCETH","binance: NAVBTC","binance: RPXBNB","binance: VETBUSD","binance: CMTETH","binance: EOSPAX","binance: ANKRUSDT","binance: ETHPAX","binance: BCHBTC","binance: BNBEUR","binance: BTGBTC","binance: CDTBTC","binance: EDOBTC","binance: POABTC","binance: SKYBTC","binance: CMTBTC","binance: NCASHETH","binance: SKYETH","binance: BNTUSDT","binance: IOTXUSDT","binance: POEBTC","binance: RLCBNB","binance: BLZBNB","binance: ZENBTC","binance: SCBTC","binance: BCHSVUSDT","binance: BTTBTC","binance: IOSTUSDT","binance: IOTABTC","binance: XRPBTC","binance: NEBLBTC","binance: VIBEETH","binance: DENTETH","binance: DASHBUSD","binance: ARPAUSDT","binance: LRCBTC","binance: AMBETH","binance: MFTBNB","binance: DCRBTC","binance: XZCXRP","binance: AIONETH","binance: IOSTBTC","binance: ZECBNB","binance: ERDUSDC","binance: MCOUSDT","binance: MANABTC","binance: VETUSDT","binance: FUNUSDT","binance: BCPTUSDC","binance: GTOTUSD","binance: BTTTRX","binance: GVTETH","binance: WAVESBNB","binance: AIONBNB","binance: BTCUSDC","binance: USDCPAX","binance: ETHBEARBUSD","binance: MODBTC","binance: DATAETH","binance: NASBNB","binance: COCOSBNB","binance: FTMUSDC","binance: FTTUSDT","binance: SNMBTC","binance: YOYOBNB","binance: BRDBTC","binance: SKYBNB","binance: ZECTUSD","binance: HSRBTC","binance: BNTBTC","binance: DOGEBTC","binance: XRPETH","binance: CVCBTC","binance: BNBBTC","binance: OMGETH","binance: GVTBTC","binance: CNDETH","binance: MITHBNB","binance: FTMBTC","binance: MBLBTC","binance: ZECBUSD","binance: AMBBNB","binance: TNBBTC","binance: STEEMBNB","binance: LOOMBNB","binance: TUSDUSDT","binance: BUSDNGN","binance: BNBTRY","binance: MBLUSDT","binance: XLMBNB","binance: OSTBNB","binance: ZILBNB","binance: EOSUSDT","binance: TFUELUSDT","binance: FETUSDT","binance: USDSBUSDS","binance: XRPBEARBUSD","binance: RDNBNB","binance: ARNBTC","binance: NASBTC","binance: RVNBTC","binance: FETBTC","binance: WINGSBTC","binance: RPXBTC","binance: ARDRBTC","binance: USDSBUSDT","binance: EOSBULLBUSD","binance: YOYOETH","binance: LSKBTC","binance: LENDETH","binance: WRXBTC","binance: DOGEUSDT","binance: DUSKBTC","binance: BTCZAR","binance: MTHETH","binance: WINGSETH","binance: TRIGETH","binance: GNTETH","binance: NEOPAX","binance: BUSDTRY","binance: NEOETH","binance: LUNETH","binance: PHBTUSD","binance: WINBTC","binance: STPTBNB","binance: XTZBTC","binance: QLCETH","binance: USDSUSDT","binance: USDSPAX","binance: ONGBTC","binance: IOSTBNB","binance: XVGETH","binance: WABIETH","binance: XRPEUR","binance: RENUSDT","binance: ETHBEARUSDT","binance: ARKETH","binance: BCHSVBTC","binance: USDCUSDT","binance: ETCPAX","binance: PERLUSDT","binance: BCHPAX","binance: NAVETH","binance: NXSBTC","binance: GOBTC","binance: BTCPAX","binance: TOMOBNB","binance: ERDBNB","binance: OGNBNB","binance: QTUMBTC","binance: BLZETH","binance: XRPUSDC","binance: TRXUSDC","binance: BTCUSDS","binance: QSPETH","binance: SYSBNB","binance: ZENETH","binance: WRXBNB","binance: ALGOBUSD","binance: BANDBTC","binance: BNBBUSD","binance: BTCBUSD","binance: MDAETH","binance: ENJETH","binance: XMRBTC","binance: STORMBNB","binance: WAVESUSDC","binance: ETHTRY","binance: BTCEUR","binance: PIVXETH","binance: BLZBTC","binance: NULSUSDT","binance: TRXPAX","binance: COCOSUSDT","binance: ARPABTC","binance: AIONUSDT","binance: NEOBNB","binance: XEMBTC","binance: LTCPAX","binance: ONEPAX","binance: ARPABNB","binance: TROYBTC","binance: STPTUSDT","binance: SNGLSBTC","binance: XMRETH","binance: XLMBTC","binance: BTCTUSD","binance: EOSBNB"}