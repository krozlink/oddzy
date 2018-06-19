package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type handler struct {
}

type requestHandler interface {
	getResponse(url string) ([]byte, error)
}

func (h *handler) getResponse(url string) ([]byte, error) {
	resp, err := http.Post(url, "", strings.NewReader(""))

	if err != nil {
		baseLog.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
