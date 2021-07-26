package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// 使用defer优化./title1/title1.go
func main() {
	// for _, url := range os.Args[1:] {
	// 	if err := title(url); err != nil {
	// 		fmt.Fprintf(os.Stderr, "title: %v\n", err)
	// 	}
	// }
	title("http://gopl.io")
}

func title(url string) error {
	// http.Get(url)获取resp.
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// resp.Header.Get获取Content-Type, "text/html;charset=utf-8"
	title := resp.Header.Get("Content-Type")
	if title != "text/html" && !strings.HasPrefix(title, "text/html;") {
		fmt.Errorf("parsing %s: not html", url)
	}

	// html.Parse(resp.Body)
	// 参考Parse的Example: example_4.go再写一遍, 详见./title2-example.go
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("html parse err: %v", err)
	}

	// forEachNode(n *html.Node, pre, post func(n *html.Node))调用
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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
