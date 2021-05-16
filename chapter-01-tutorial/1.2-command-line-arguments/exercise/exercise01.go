package main

import (
	"fmt"
	"os"
	"strings"
)

// print the name of command, the name of os.Args[0].
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
