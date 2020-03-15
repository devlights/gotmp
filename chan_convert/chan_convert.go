package main

import (
	"fmt"
)

func main() {
	type (
		doneCh chan struct{}
	)

	done := make(chan struct{})
	done2 := make(doneCh)

	var done3 doneCh
	done3 = done

	var done4 doneCh
	done4 = done2

	close(done3)
	close(done4)

	for _, v := range []doneCh{done, done2, done3, done4} {
		ch := v
		_, ok := <-ch
		fmt.Println(ok)
	}
}
