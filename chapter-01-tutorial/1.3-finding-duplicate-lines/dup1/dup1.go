package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appears more than once in the standard input, preceded by its count.
func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	// Note: ignore the potential errors from input.Err()
	for input.Scan() {
		// break the dead for-loop.
		if input.Text() == ":wq" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
