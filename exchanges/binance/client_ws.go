package binance

import (
	"context"
	"errors"
	"io/ioutil"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	baseStreamUrl = "wss://stream.binance.com:9443/ws/"
)

func connectWs(ctx context.Context, urlPath string, processMessage func([]byte)) error {
	log.WithField("url", baseStreamUrl+urlPath).Trace("websocket connecting")
	c, resp, err := websocket.DefaultDialer.Dial(
		baseStreamUrl+urlPath,
		nil,
	)
	if err != nil {
		log.WithError(err).Error("ws dial error")
		return err
	}

	if resp.StatusCode != 101 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithError(err).Error("read body error")
		}
		defer resp.Body.Close()

		log.WithFields(logrus.Fields{
			"code": resp.StatusCode,
			"body": string(bodyBytes),
		}).Error("connected, but not websocket?")
	}
	log.Trace("websocket connected")

	defer func() {
		log.Debug("ws goroutine close")
		err := c.Close()
		if err != nil {
			log.WithError(err).Debug("ws close error")
		}
	}()

	timeoutTicker := time.NewTicker(time.Minute)
	lastMsgTime := time.Now()
	defer timeoutTicker.Stop()

	for {
		select {
		case now := <-timeoutTicker.C: //TODO need diffrent gorutine?
			if now.Sub(lastMsgTime) > 5*time.Minute {
				log.WithError(err).Error("ws timeout")
				return errors.New("ws timeout")
			}
		case <-ctx.Done():
			log.Debug("ws resived ctx.Done")
			return context.Canceled
		default:
			startWaitMessageTime := time.Now()
			messageType, message, err := c.ReadMessage()
			if err != nil {
				log.WithError(err).Error("ws read msg error")
				return errors.New("ws read msg error")
			}
			if messageType != websocket.TextMessage {
				log.WithField("messageType", messageType).Debug("read not text message")
				continue
			}
			lastMsgTime = time.Now()
			websocketWaitNewMessageTime.Add(float64(lastMsgTime.Sub(startWaitMessageTime)) / 1e9)

			processMessage(message)
			websocketMessageProcessingTime.Add(float64(time.Now().Sub(lastMsgTime)) / 1e9)
		}
	}
}
