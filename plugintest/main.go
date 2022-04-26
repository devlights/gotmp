package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("lib/lib.so")
	if err != nil {
		panic(err)
	}

	sym, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}

	sym.(func())()

	sym, err = p.Lookup("I")
	if err != nil {
		panic(err)
	}

	fmt.Println(sym.(*int))
	fmt.Println(*(sym.(*int)))
}
