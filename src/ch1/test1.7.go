package main
/*
	函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，
避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
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
