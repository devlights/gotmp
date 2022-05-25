package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func add[E constraints.Signed](x, y E) E {
	return x + y
}

func main() {
	fmt.Println(add(int32(1), int32(3)))
	fmt.Println(add(int8(1), int8(2)))
}
