package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fetch("http://gopl.io")
}

// fetch下载URL并返回本地文件的名字和长度.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, nil
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	file, err := os.Create(local)
	if err != nil {
		return "", 0, nil
	}
	n, err = io.Copy(file, resp.Body)
	// 关闭文件, 并保留错误信息.
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}
