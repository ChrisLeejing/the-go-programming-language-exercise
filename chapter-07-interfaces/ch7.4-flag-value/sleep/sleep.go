package main

import (
	"flag"
	"fmt"
	"time"
)

var duration = flag.Duration("period", 5*time.Second, "sleep 5s")

func main() {
	flag.Parse() // 用于解析command-line输入的参数解析.
	fmt.Printf("sleep for %v\n", *duration)
	time.Sleep(*duration)
	fmt.Println(111)
	// $ go build -o sleep sleep.go
	//
	// Chris@LAPTOP-ES8KBP40 MINGW64 /e/Code/Go/go_books_resource_code/the-go-programming-language-exercise/chapter-07-interfaces/ch7.4-flag-value/sleep (test)
	// $ ./sleep
	// sleep for 5s
	// 111
	//
	// Chris@LAPTOP-ES8KBP40 MINGW64 /e/Code/Go/go_books_resource_code/the-go-programming-language-exercise/chapter-07-interfaces/ch7.4-flag-value/sleep (test)
	// $ ./sleep -period 1s
	// sleep for 1s
	// 111
	//
	// Chris@LAPTOP-ES8KBP40 MINGW64 /e/Code/Go/go_books_resource_code/the-go-programming-language-exercise/chapter-07-interfaces/ch7.4-flag-value/sleep (test)
	// $ ./sleep -period 1m10s
	// sleep for 1m10s
	// 111
	//
	// Chris@LAPTOP-ES8KBP40 MINGW64 /e/Code/Go/go_books_resource_code/the-go-programming-language-exercise/chapter-07-interfaces/ch7.4-flag-value/sleep (test)
	// $ ./sleep -period "1 hour"
	// invalid value "1 hour" for flag -period: parse error
	// Usage of E:\Code\Go\go_books_resource_code\the-go-programming-language-exercise\chapter-07-interfaces\ch7.4-flag-value\sleep\sleep:
	//  -period duration
	//        sleep 5s (default 5s)
}
