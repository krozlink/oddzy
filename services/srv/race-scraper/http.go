package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct {
}

type requestHandler interface {
	getResponse(url string) ([]byte, error)
}

func (h *handler) getResponse(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
