package main

import (
	"fmt"
	"regexp"
)

func main() {
	httpSchemeRE := MustCompile("^https?:")
	fmt.Println("httpSchemeRE: ", httpSchemeRE)
}

// func Compile(expr string) (*regexp.Regexp, error) {
// 	/*
// 		处理逻辑
// 	*/
// 	return nil, nil
// }

func MustCompile(expr string) *regexp.Regexp {
	re, err := regexp.Compile(expr)
	if err != nil {
		panic(err)
	}

	return re
}
