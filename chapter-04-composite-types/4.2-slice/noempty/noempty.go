package main

import "fmt"

func main() {
	fmt.Printf("%q\n", noEmpty([]string{" ", "a", "b", " "}))
	fmt.Printf("%q\n", noEmpty2([]string{" ", "a", "b", " "}))
	fmt.Printf("%q\n", noEmpty3([]string{" ", "a", "b", " "}))
}

// coirlee
func noEmpty(s []string) []string {
	var res []string
	for _, v := range s {
		if v != " " {
			res = append(res, v)
		}
	}

	return res
}

// Author: Alan or Brian
func noEmpty2(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != " " {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func noEmpty3(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != " " {
			out = append(out, s)
		}
	}

	return out
}
