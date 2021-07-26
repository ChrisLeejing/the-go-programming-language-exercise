package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a unsafe.Pointer
	fmt.Printf("%T\n", a)
}
