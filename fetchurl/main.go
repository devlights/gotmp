package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	errCh := make(chan error, 2)
	url := "https://www.google.co.jp/"
	
	ch1 := make(chan *http.Response, 1)
	go func() {
		defer close(ch1)

		resp, err := http.Get(url)
		if err != nil {
			errCh <- err
		}

		ch1 <- resp
	}()

	ch2 := make(chan int)
	go func() {
		defer close(ch2)

		bytes, err := ioutil.ReadAll((<-ch1).Body)
		if err != nil {
			errCh <- err
		}

		ch2 <- len(bytes)
	}()

	log.Printf("[fetch-url] %d bytes recv", <-ch2)

	close(errCh)
	for e := range errCh {
		log.Print(e)
	}

	return 0
}