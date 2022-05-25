// Goのスライスで s[a:b:c] と指定した場合の 3つ目 の値指定について
//
// REFERENCES
//   - https://stackoverflow.com/questions/27938177/golang-slice-slicing-a-slice-with-sliceabc
//   - https://stackoverflow.com/questions/12768744/re-slicing-slices-in-golang/18911267#18911267
//   - https://tip.golang.org/doc/go1.2#three_index
//   - https://go.dev/ref/spec#Slice_expressions
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
