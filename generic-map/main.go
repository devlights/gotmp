package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	var (
		m1 = map[string]int{ "hello": 100, "world": 101 }
		m2 = map[string]int{ "world": 101, "hello": 100 }
		k []string
		v []int
	)

	k = maps.Keys(m1)
	v = maps.Values(m1)

	fmt.Println(k, v)
	fmt.Println(maps.Equal(m1, m2))

	maps.Clear(m1)
	fmt.Println(m1, m2)
}