package service

import (
	"net"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func NewDocumentWithTimeout(url string, timeout time.Duration) (*goquery.Document, error) {
	// Load the URL
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: timeout,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}
	client := &http.Client{
		Timeout:   timeout,
		Transport: netTransport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		return nil, e
	}

	res, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	return goquery.NewDocumentFromResponse(res)
}
