package main
/*
	 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
*/

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		response, err := http.Get(url)
		if (err != nil) {
			fmt.Println("fetch", url, err)
			os.Exit(1)
		}

		fmt.Println(response.Status)
		response.Body.Close()
	}
}
