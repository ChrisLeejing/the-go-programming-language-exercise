package main

import "fmt"

func main() {

}

type Syntax struct {
}

func Parse(input string) (s *Syntax, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internel err: %v", p)
		}
	}()

	// 处理逻辑.
	return nil, nil
}
