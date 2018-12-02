package main
/*
	修改echo程序，使其打印value和index，每个value和index显示一行
*/
import (
	"fmt"
	"os"
)

func main(){
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
}