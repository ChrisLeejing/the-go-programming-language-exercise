package main

import "fmt"

func main() {
	a := []string{"1", "2", "3", "4", "0"}
	b := []string{"0", "2", "3", "4", "5"}
	fmt.Println(equal(a, b))
}

// coirlee
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Author: Alan or Brian
func equal2(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
