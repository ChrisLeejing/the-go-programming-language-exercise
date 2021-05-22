package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch url: %s, err: %v", url, err)
			os.Exit(1)
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		statusCode := resp.StatusCode
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading url: %s, err: %v", url, err)
			os.Exit(1)
		}

		fmt.Printf("fetch reading url: %v\nstatus code: %d", string(bytes), statusCode)
	}
}
