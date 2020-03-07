package main

import (
	"fmt"
	"os"
)

type (
	command interface {
		Run()
	}

	cmdA struct {
		data string
	}
)

var _ command = (*cmdA)(nil)
var _ fmt.Stringer = (*cmdA)(nil)

func newCmdA() command {
	c := new(cmdA)
	c.data = "helloworld"
	return c
}

func (c *cmdA) Run() {
	c.data = "worldhello"
}

func (c *cmdA) String() string {
	return c.data
}

func main() {
	os.Exit(run())
}

func run() int {
	c := newCmdA()
	c.Run()
	fmt.Println(c)
	return 0
}
