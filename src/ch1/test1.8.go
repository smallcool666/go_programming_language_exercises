package main
/*
	修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
*/

import (
	"fmt"
	"io"
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

		num, err := io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if (err != nil) {
			fmt.Println("copy reponse", url, err)
		}
		//输出copy成功的字节数
		fmt.Println(num)
	}
}
