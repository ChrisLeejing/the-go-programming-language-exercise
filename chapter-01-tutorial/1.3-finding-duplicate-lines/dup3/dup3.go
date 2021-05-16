package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
