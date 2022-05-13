package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/output"
)

const (
	SUCCESS = iota
	FAILURE
)

func main() {
	os.Exit(run())
}

func run() int {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 10*time.Millisecond)
		done             = procCtx.Done()
	)
	defer mainCxl()
	defer procCxl()

	var (
		gen    = chans.Generator(done, 1, 2, 3, 4, 5)
		proc   = chans.Map(done, gen, func(v int) int { return v * 2 })
		convFn = func(v *chans.MapValue[int]) string { return fmt.Sprintf("%v --> %v", v.Before, v.After) }
		conv   = chans.Convert(done, proc, convFn)
		chunks = chans.Chunk(done, conv, 3)
	)

	for v := range chunks {
		output.Stdoutl("[len]", len(v))
		for _, x := range v {
			output.Stdoutf("[===>]", "\t%v\n", x)
		}
	}

	return SUCCESS
}
