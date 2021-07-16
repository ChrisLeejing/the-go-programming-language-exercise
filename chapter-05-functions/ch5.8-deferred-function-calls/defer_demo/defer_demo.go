package main

import "sync"

func main() {

}

// defer use for sync.Mutex.
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
