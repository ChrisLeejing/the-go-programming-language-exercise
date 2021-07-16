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
		if err := title3(url); err != nil {
			fmt.Fprintf(os.Stderr, "title err: %v\n", err)
		}
	}
}

// 参考: html.Parse(resp.Body) example示例程序进行编写.
func title3(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	title := resp.Header.Get("Content-Type")
	if title != "text/html" && !strings.HasPrefix(url, "text/html;") {
		fmt.Errorf("title: %s not html", url)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("parsing err: %v", err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil
}
