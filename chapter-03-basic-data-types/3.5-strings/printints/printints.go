package main

import (
	"bytes"
	"fmt"
	"time"
)

// time coirlee:  506300  Nanoseconds.
// [11, 22, 33, 44, 55, 66, 77, 88, 99]
// time Alan or Brian:  0  Nanoseconds.
// [11, 22, 33, 44, 55, 66, 77, 88, 99]
func main() {
	result1 := intsToString([]int{11, 22, 33, 44, 55, 66, 77, 88, 99})
	fmt.Println(result1)
	result2 := intsToString2([]int{11, 22, 33, 44, 55, 66, 77, 88, 99})
	fmt.Println(result2)
}

// [1, 2, 3] -> [1, 2, 3] string.
func intsToString(values []int) string {
	t1 := time.Now().Nanosecond()
	var buf bytes.Buffer
	begin := "["

	for i, value := range values {
		if i > 0 {
			buf.WriteByte(',')
			// buf.WriteRune('中')
			// buf.WriteByte('中') // .\printints.go:29:18: constant 20013 overflows byte
			buf.WriteByte(' ')
		}
		buf.WriteString(fmt.Sprintf("%d", value))
	}

	end := "]"
	s := begin + buf.String() + end
	t2 := time.Now().Nanosecond()
	fmt.Println("time coirlee: ", t2-t1, " Nanoseconds.")
	return s
}

func intsToString2(values []int) string {
	t1 := time.Now().Nanosecond()
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, value := range values {
		if i > 0 {
			buf.WriteByte(',')
			// buf.WriteRune('中')
			// buf.WriteByte('中') // .\printints.go:29:18: constant 20013 overflows byte
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", value)
	}

	buf.WriteByte(']')
	t2 := time.Now().Nanosecond()
	fmt.Println("time Alan or Brian: ", t2-t1, " Nanoseconds.")

	return buf.String()
}
