package main

import (
	"bufio"
	"fmt"
	"os"
)

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
func main() {
	counts := make(map[string]FileData)
	files := os.Args[1:]
	if len(files) == 0 {
		// stdin
		countLine(os.Stdin, counts)
	} else {
		// file input
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Modify dup2: %v\n", err)
				continue
			}
			countLine(f, counts)
			// close file
			f.Close()
		}
	}
	// print
	for line, fileData := range counts {
		n := fileData.lineCount
		fileName := fileData.lineFileName
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileName)
		}
	}
}

type FileData struct {
	lineCount    int
	lineFileName string
}

func countLine(file *os.File, counts map[string]FileData) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if input.Text() == ":wq" {
			break
		}
		fileData := counts[input.Text()]
		fileData.lineCount++
		fileData.lineFileName = file.Name()
		// debug note: remember to reset counts.
		counts[input.Text()] = fileData
	}
}
