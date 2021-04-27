package signals_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/devlights/gotmp/signals"
)

func ExampleReady2() {
	var (
		wg sync.WaitGroup
	)

	fn := func() (ready *signals.Ready2, values <-chan int) {
		r := signals.NewReady2()
		ch := make(chan int)

		go func() {
			defer close(ch)

			time.Sleep(2 * time.Second)
			ready.Signal()

			for _, v := range []int{1, 2, 3, 4, 5} {
				ch <- v
			}
		}()

		return r, ch
	}

	ready, values := fn()

	for _, v := range []int{1, 2, 3} {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Printf("[%d]:begin wait\n", i)
			ready.Wait()
			fmt.Printf("[%d]:end wait\n", i)

			for v := range values {
				fmt.Printf("%d\n", v)
			}
		}(v)
	}

	wg.Wait()

	// Unordered output:
	// [1]:begin wait
	// [2]:begin wait
	// [3]:begin wait
	// [1]:end wait
	// [2]:end wait
	// [3]:end wait
	// 1
	// 2
	// 3
	// 4
	// 5
}
