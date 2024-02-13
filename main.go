package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	brasilApiResult := make(chan []byte)
	viacepResult := make(chan []byte)
	cep := 22061030

	// Consuming APIs
	go func() {
		log.Println("Requesting BrasilApi ...")
		res, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%d", cep))
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		log.Println("... BrasilApi responded")

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		brasilApiResult <- body
	}()

	go func() {
		log.Println("Requesting ViaCep ...")
		res, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%d/json", cep))
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		log.Println("... BrasilApi responded")

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		viacepResult <- body
	}()

	// Consumes the fastest response
	select {
	case msg := <-brasilApiResult:
		log.Println("Result from BrasilAPI: ", string(msg))

	case msg := <-viacepResult:
		log.Println("Result from ViaCep: ", string(msg))

	case <-time.After(time.Second * 1):
		println("Request Timeout")
	}
}
