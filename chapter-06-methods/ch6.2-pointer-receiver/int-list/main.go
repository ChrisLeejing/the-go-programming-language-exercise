package main

func main() {

}

// IntList是int型链表
// IntList为nil时, 表示空列表
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum 返回列表元素的总和
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Value
}
