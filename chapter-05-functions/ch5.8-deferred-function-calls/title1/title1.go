package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := title(url); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

// 获取一个HTML文档, 然后输出它的标题. title函数检测从服务器端返回的Content-Type头部, 如果不是HTML, 则返回错误.
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Content-Type ("text/html;charset=utf-8")
	title := resp.Header.Get("Content-Type")
	if title != "text/html" && !strings.HasPrefix(title, "text/html;") {
		resp.Body.Close()
		fmt.Errorf("%s has type %s, not text/html", url, title)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Errorf("parsing %s as html: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

// .\ch5.5-function-values\outline2\outline2.go
func forEachNode(n *html.Node, pre, post func(*html.Node)) {
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
