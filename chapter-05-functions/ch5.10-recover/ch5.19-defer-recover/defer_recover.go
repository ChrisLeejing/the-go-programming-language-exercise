package main

import "fmt"

func main() {
	fmt.Println(weird())
}

func weird() (result string) {
	defer func() {
		// if err := recover(); err != nil {
		// 	fmt.Println(err)
		// }
		recover()
		result = "hi"
	}()

	panic("oh, my god.")
}
