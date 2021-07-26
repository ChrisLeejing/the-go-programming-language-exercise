package main

import "sync"

func main() {

}

// 版本1:
var (
	mu      sync.Mutex // 保护mapping
	mapping = make(map[string]string)
)

func LookUp(key string) string {
	mu.Lock()
	value := mapping[key]
	mu.Unlock()

	return value
}

// 版本2: 变量名更加贴切, 且sync.Mutex是内嵌的.
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func LookUp2(key string) string {
	cache.Lock()
	v := mapping[key]
	cache.Unlock()

	return v
}
