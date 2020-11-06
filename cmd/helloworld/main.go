package main

import (
	"fmt"
	"os"
	"sync"
)

const (
	message = "helloworld"
)

func main() {
	os.Exit(run())
}

func run() int {
	var (
		wg    = sync.WaitGroup{}
		outCh = make(chan rune, len(message))
	)

	// start input goroutines
	wg.Add(len(message))
	for _, c := range message {
		go func(c rune, ch chan<- rune) {
			defer wg.Done()
			ch <- c
		}(c, outCh)
	}

	// start closer goroutine
	go func() {
		wg.Wait()
		close(outCh)
	}()

	// start output goroutine
	func(ch <-chan rune) {
		for c := range ch {
			fmt.Printf("%c", c)
		}

		fmt.Println("")
	}(outCh)

	return 0
}
