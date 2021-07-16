package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// bigSlowOperation()
	// _ = double(4)
	// fmt.Println(triple(4))

}

func bigSlowOperation() {
	// 注意尾部()加与不加语句执行的顺序.
	// defer trace("bigSlowOperation")
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() {
		log.Printf("exit %s (%v)", msg+"++", time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
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
		if err := doFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func doFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}
