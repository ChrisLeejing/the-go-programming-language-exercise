package main

import "fmt"

func main() {
	var c ByteCounter
	n, _ := c.Write([]byte("hello"))
	fmt.Printf("len(c) = %d\n", n)

	c = 0
	name := "chris"
	// Fprintf()内部调用Write().
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p)) // 转换int类型为ByteCounter
	return len(p), nil
}
