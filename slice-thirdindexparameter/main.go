package main

import "fmt"

func main() {
	s := []string{
		"golang",
		"dotnet",
		"java",
		"python",
		"ruby",
	}

	s2 := s[1:3]
	// len = high(2nd) - low(1st) = 3 - 1 = 2
	// cap = max(3rd) - low(1st) = len(s) - 1 = 5 - 1 = 4
	fmt.Println(s2, len(s2), cap(s2))

	s3 := s[1:3:4]
	// len = high(2nd) - low(1st) = 3 - 1 = 2
	// cap = max(3rd) - low(1st) = 4 - 1 = 3
	fmt.Println(s3, len(s3), cap(s3))
}
