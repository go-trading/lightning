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
	HttpServer         *http.Server
	HttpServerExitDone *sync.WaitGroup
	Node               *core.Node
}

func (w *WebService) Init(node *core.Node, config *core.ServiceConfig) error {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	w.HttpServer = &http.Server{Addr: ":8080"}
	w.HttpServerExitDone = &sync.WaitGroup{}
	w.Node = node

	return nil
}

func (w *WebService) Start() error {
	w.HttpServerExitDone.Add(1)
	go func() {
		defer w.HttpServerExitDone.Done() // let main know we are done cleaning up

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

	w.HttpServerExitDone.Wait()
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
