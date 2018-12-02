package main

import (
	"fmt"
	"io/ioutil"
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
		body, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if (err != nil) {
			fmt.Println("fetch:reading", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", body)
	}
}
