package main

import "fmt"

func main() {
	var runes []rune

	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	// go_1.16.3\go\src\fmt\doc.go
	// %q	a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%q", runes)
}
