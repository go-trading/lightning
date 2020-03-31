package binance

import (
	"io/ioutil"
	"net/http"
)

const (
	baseUrl = "https://api.binance.com"
)

func (b *Binance) clientGet(urlPath string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(baseUrl + urlPath)
	if err != nil {
		log.WithError(err).WithField("urlPath", urlPath).Error("can't get")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).WithField("urlPath", urlPath).Error("can't read body")
		return nil, err
	}
	return body, nil
}
