package web

import (
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/go-trading/lightning/core"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("p", "web")
)

func Register(availableServices core.AvailableServices) {
	availableServices.Register("web", &WebService{})
}

type WebService struct {
	sync.WaitGroup
	HttpServer *http.Server
	Node       *core.Node
}

func (w *WebService) Init(node *core.Node, config *core.ServiceConfig) error {
	w.Node = node

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	w.HttpServer = &http.Server{Addr: config.GetString("address")}

	if err := config.GetErrors(); err != nil {
		log.WithError(err).Error("maybe incorrect web config")
		return err
	}

	return nil
}

func (w *WebService) Start() error {
	w.Add(1)
	go func() {
		defer w.Done() // let main know we are done cleaning up

		if err := w.HttpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return nil
}

func (w *WebService) Stop() error {
	if err := w.HttpServer.Shutdown(context.TODO()); err != nil {
		log.WithError(err).Error("Can't shutdown http server")
		return err
	}

	w.Wait()
	return nil
}

func (w *WebService) Name() string {
	return "WebGUI"
}

func (w *WebService) Status() core.ServiceStatus {
	return core.Stopped
}

func (w *WebService) Error() error {
	return nil
}

func (w *WebService) SubscribeStatus(func(core.ServiceStatus))   {}
func (w *WebService) UnsubscribeStatus(func(core.ServiceStatus)) {}
