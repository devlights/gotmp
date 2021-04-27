package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/devlights/gotmp/signals"
)

func main() {
	os.Exit(run())
}

func run() int {
	// ------------------------------
	// ↓↓↓ PUT SAMPLE CODE BELOW ↓↓↓
	// ------------------------------
	ready1()
	fmt.Println("-------------------------------------------")
	ready2()

	return 0
}

func ready2() {
	var (
		wg = sync.WaitGroup{}
	)

	ready, ch := asyncfn2()
	for _, i := range []int{1, 2, 3, 4, 5} {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Printf("[%d]:wait ready...\n", i)
			ready.Wait()
			fmt.Printf("[%d]:wait done \n", i)

			for v := range ch {
				fmt.Printf("[%d]: value:%d\n", i, v)
				time.Sleep(1 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}

func ready1() {
	var (
		wg = sync.WaitGroup{}
	)

	ready, ch := asyncfn1()
	for _, i := range []int{1, 2, 3, 4, 5} {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Printf("[%d]:wait ready...\n", i)
			ready.Wait()
			fmt.Printf("[%d]:wait done \n", i)

			for v := range ch {
				fmt.Printf("[%d]: value:%d\n", i, v)
				time.Sleep(1 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}

func asyncfn2() (*signals.Ready2, <-chan int) {
	var (
		r = signals.NewReady2()
		c = make(chan int)
	)

	go func() {
		defer close(c)

		println("begin sleep")
		time.Sleep(2 * time.Second)
		println("end sleep")

		println("signal ready")
		r.Signal()

		for _, v := range []int{1, 2, 3, 4, 5} {
			c <- v
		}
	}()

	return r, c
}


func asyncfn1() (*signals.Ready1, <-chan int) {
	var (
		r = signals.NewReady1()
		c = make(chan int)
	)

	go func() {
		defer close(c)

		println("begin sleep")
		time.Sleep(2 * time.Second)
		println("end sleep")

		println("signal ready")
		r.Signal()

		for _, v := range []int{1, 2, 3, 4, 5} {
			c <- v
		}
	}()

	return r, c
}
