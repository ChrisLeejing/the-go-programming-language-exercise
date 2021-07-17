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
			fmt.Fprintf(os.Stderr, "title err: %v\n", err)
		}
	}
}

// 返回文档中第一个非空标题元素; 如果标题不唯一, 则返回错误.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic.
		case bailout{}:
			// 预期的错误信息, 进行返回.
			err = fmt.Errorf("multiple title elements")
		default:
			// 非预期错误, 进行进行panic.
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements. <title>a</title> <title>b</title>(第二次title会抛panic.)
			}

			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}

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

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check Content-Type is html, eg: "text/html;charset=utf-8"
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not html", url, ct)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML %v", url, err)
	}
	title, err := soleTitle(node)
	if err != nil {
		return err
	}
	fmt.Printf("title: %s\n", title)

	return nil
}
