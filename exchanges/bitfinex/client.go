package bitfinex

import (
	"io/ioutil"
	"net/http"
)

const (
	basePubUrl  = "https://api-pub.bitfinex.com"
	baseAuthUrl = "https://api.bitfinex.com"
)

func (b *Bitfinex) clientGetPub(urlPath string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(basePubUrl + urlPath)
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

func (b *Bitfinex) clientGetAuth(urlPath string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(baseAuthUrl + urlPath)
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
