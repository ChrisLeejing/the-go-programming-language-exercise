package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(s)
	// fmt.Println("coirlee: ", remove(s, 2))
	// fmt.Println("Alan or Brian: ", remove2(s, 2))

	fmt.Println(s)
	fmt.Println("Alan or Brian: ", remove3(s, 2))
}

// coirlee
func remove(ints []int, i int) []int {
	var values []int
	for k, v := range ints {
		if k == i {
			continue
		}

		values = append(values, v)
	}

	return values
}

// Author: Alan or Brian
// need order.
func remove2(ints []int, i int) []int {
	copy(ints[i:], ints[i+1:])
	return ints[:len(ints)-1]
}

// no need order.
func remove3(ints []int, i int) []int {
	ints[i] = ints[len(ints)-1]
	return ints[:len(ints)-1]
}
