package main

import "fmt"

func main() {
	m := map[string]int{
		"hoge": 100,
		"fuga": 200,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
