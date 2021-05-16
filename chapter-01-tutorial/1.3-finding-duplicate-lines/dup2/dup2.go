package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup2 prints the countLine and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		// stdin
		countLine(os.Stdin, counts)
	} else {
		// file input
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLine(f, counts)

			// close file
			f.Close()
		}
	}

	// print
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLine(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if input.Text() == ":wq" {
			break
		}
		counts[input.Text()]++
	}
}
