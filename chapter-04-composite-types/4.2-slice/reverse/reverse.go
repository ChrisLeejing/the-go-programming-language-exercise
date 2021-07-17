package main

import "fmt"

func main() {
	values := []int{0, 1, 2, 3, 4, 5}
	// reverse(values[:2])
	// reverse(values[2:])
	// reverse(values)
	reverse2(values[:2])
	reverse2(values[2:])
	reverse2(values)

	fmt.Println(values)

}

// coirlee
func reverse(values []int) []int {
	for i, j := 0, len(values)-1; i <= (len(values)-1)/2; {
		values[i], values[j] = values[j], values[i]
		i++
		j--
	}

	return values
}

// Author: Alan or Brian
func reverse2(s []int) []int {
	// i, j = i+1, j-1 -> can not be i, j = i++, j++
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
