package main

import (
	"fmt"
)

// WithBeginEnd は、指定された関数の呼び出し前後でログを出力して実行します
func WithBeginEnd(message string, f func()) {
	fmt.Printf("======= START [%s] =======\n", message)
	f()
	fmt.Printf("======= END   [%s] =======\n", message)
}