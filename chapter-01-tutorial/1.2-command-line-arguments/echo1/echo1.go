package main

import (
	"fmt"
	"os"
)

// echo1.go print the command line arguments.
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// $ ./echo1.exe a b c 1 2 34       jd       j           hello       你好
	// a b c 1 2 34 jd j hello 你好
}
