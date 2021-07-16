package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	bigSlowOperation()
	_ = double(4)
	fmt.Println(triple(4))
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s.\n", msg)
	return func() {
		fmt.Printf("exit %s %v\n", msg, time.Since(start))
	}
}

// error example.
func openFiles(filenames []string) error {
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		defer file.Close() // Possible resource leak, 'defer' is called in a 'for' loop
	}

	return nil
}

// correct example.
func openFiles2(filenames []string) error {
	for _, filename := range filenames {
		if err := openFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func openFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 处理下面的逻辑.
	return nil
}
