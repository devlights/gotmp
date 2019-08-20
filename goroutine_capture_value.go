package main

import (
	"fmt"
	"sync"
)

// GoRoutineCaptureValue は、Goroutineでループの値は局所変数にとって利用しないといけないサンプル
// REFERENCES:: http://bit.ly/2HdAN1m
func GoRoutineCaptureValue() {
	runConcurrentWrong(func1, func2)
	runConcurrentCorrect(func1, func2)
}

func runConcurrentCorrect(funcs ...func()) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for _, fn := range funcs {
		wg.Add(1)

		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}
}

func runConcurrentWrong(funcs ...func()) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for _, fn := range funcs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fn()
		}()
	}
}

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}
