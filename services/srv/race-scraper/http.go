package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type httpHandler struct {
	attempts int
	interval int
}

type requestHandler interface {
	getResponse(url string) ([]byte, error)
}

func newHTTPHandler() *httpHandler {
	return &httpHandler{
		attempts: 3,
		interval: 1000,
	}
}

func (h *httpHandler) getResponse(url string) ([]byte, error) {
	for i := 1; i <= h.attempts; i++ {
		resp, err := http.Post(url, "", strings.NewReader(""))
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		if err != nil {
			if i == h.attempts {
				baseLog.Fatalf("max %v attempts exceeded in requesting url %v - %v", i, url, err)
			}

			baseLog.Warnf("attempt %v of %v when requesting url %v - %v", i, h.attempts, url, err)
			time.Sleep(time.Millisecond * time.Duration(h.interval))
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		return body, err
	}

	return nil, nil
}
