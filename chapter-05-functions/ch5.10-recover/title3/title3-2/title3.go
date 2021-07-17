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

func title(url string) error {
	// get resp.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check Content-Type is HTML, eg: "text/html;charset=utf-8"
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		fmt.Errorf("%s type %s, not html\n", url, ct)
	}

	// parse html.
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML, err: %v\n", url, err)
	}
	title, err := soleTitle(doc)
	if err != nil {
		return err
	}
	fmt.Println("title:", title)

	return nil
}

func soleTitle(doc *html.Node) (title string, err error) {
	// 定义bailout struct表示多个title.
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic.
		case bailout{}:
			// expected err.
			err = fmt.Errorf("multiple title elements")
		default:
			// unexpected err, carry on panicking.
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements.
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
