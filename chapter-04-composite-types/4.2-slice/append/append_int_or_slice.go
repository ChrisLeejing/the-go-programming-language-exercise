package main

import (
	"fmt"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d, cap=%d, %v\n", i, cap(y), y)
		x = y
		y = append(y, i)
	}

	fmt.Println("-------------------")

	var s1 []int
	s1 = appendSlice(s1, 0, 1, 2)
	fmt.Printf("1 cap=%d, %v\n", cap(s1), s1)
	s1 = appendSlice(s1, 3, 4, 5)
	fmt.Printf("2 cap=%d, %v\n", cap(s1), s1)
	s1 = appendSlice(s1, 6, 7, 8)
	fmt.Printf("3 cap=%d, %v\n", cap(s1), s1)
	s1 = appendSlice(s1, 9, 10, 11)
	fmt.Printf("4 cap=%d, %v\n", cap(s1), s1)

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// x 还有剩余空间, 可以进行赋值.
		z = x[:zlen]
	} else {
		// x 没有剩余空间, 需要进行扩容.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x) // zcap = len(x) + 1 < 2*len(x) => len(x) > 1.
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendSlice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		// x 还有剩余空间, 可以进行赋值.
		z = x[:len(x)]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}
