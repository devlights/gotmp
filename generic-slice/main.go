package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var (
		s1 = []string{"hello", "world"}
		s2 = []string{"hello", "world"}
		s3 = []string{"world", "hello"}
		s4 = []int{100, 101}
	)

	fmt.Println(slices.Equal(s1, s2))
	fmt.Println(slices.Equal(s2, s3))
	// compile error
	//fmt.Println(slices.Equal(s1, s4))

	s5 := slices.Insert(s4, 1, 999)
	fmt.Println(s4, s5)

	idx := slices.Index(s5, 999)
	s6 := slices.Delete(s5, idx, idx+1)
	fmt.Println(s5, s6)
}
