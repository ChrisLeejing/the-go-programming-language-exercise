package main

import "golang.org/x/net/html"

func main() {

}

// forEachNode调用pre(x)和post(x)遍历以n为根的树中每个节点x.
// 两个函数是可选的.
// pre 在子节点被访问前(前序)调用, post在被访问后(后序)调用.
func forEachNode(n *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
