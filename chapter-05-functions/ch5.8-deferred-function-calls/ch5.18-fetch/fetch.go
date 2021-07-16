package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		filename, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s err: %v\n", url, err)
		}
		fmt.Fprintf(os.Stdout, "%s => %s (%d bytes)\n", url, filename, n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, nil
	}
	defer resp.Body.Close()

	locale := path.Base(resp.Request.URL.Path)
	if locale == "/" {
		locale = "index.html"
	}
	file, err := os.Create(locale)
	if err != nil {
		return "", 0, nil
	}

	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(file, resp.Body)

	return locale, n, err
}
