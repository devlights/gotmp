package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	delimiter = []byte("[TAB]")
)

func scanTab(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, delimiter); i >= 0 {
		return i + len(delimiter), data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func main() {
	str := "hello[TAB]world[TAB]golang[TAB]javascript"
	sca := bufio.NewScanner(strings.NewReader(str))

	//sca.Split(bufio.ScanWords)
	sca.Split(scanTab)

	for sca.Scan() {
		fmt.Fprintln(os.Stdout, sca.Text())
	}
}
