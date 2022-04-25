package main

import (
	"context"
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
	// ------------------------------
	// ↓↓↓ PUT SAMPLE CODE BELOW ↓↓↓
	// ------------------------------
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
	)
	defer mainCxl()

	var (
		procCtx, procCxl = context.WithTimeout(mainCtx, 1*time.Second)
		done             = procCtx.Done()
		gen              = chans.Generator(done, 1, 2, 3, 4, 5)
		proc             = chans.Map(done, gen, func(v int) int { return v * 2 })
		out              = chans.Buffer(done, proc, 3)
	)
	defer procCxl()

	for v := range out {
		output.Stdoutl("[len]", len(v))
		for _, x := range v {
			output.Stdoutf("[===>]", "\t%v\n", x.After)
		}
	}

	return SUCCESS
}
