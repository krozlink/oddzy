package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

type oddsResponse struct {
	r string
}

func getResponse(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	result := decode(body)
	return result
}

func decode(response []byte) []byte {

	rot := make([]byte, len(response))
	result := make([]byte, len(response))
	for i, b := range response {
		rot[i] = b
	}
	base64.StdEncoding.Decode(result, rot)
	return result
}

func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}
