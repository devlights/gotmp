package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(homeDir)
}
