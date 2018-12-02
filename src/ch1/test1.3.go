package main
/*
	上手实践前面提到的strings.Join和直接Println，并观察输出结果的区别
*/
import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//C:\Users\aaa\AppData\Local\Temp\go-build022935074\b001\exe\test1.3.exe 1 2 3
	fmt.Println(strings.Join(os.Args, " "))
	//[C:\Users\aaa\AppData\Local\Temp\go-build022935074\b001\exe\test1.3.exe 1 2 3]
	fmt.Println(os.Args)
}