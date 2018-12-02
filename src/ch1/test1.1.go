package main
/*
	修改echo程序，使其能够打印os.Args[0]
*/
import (
	"fmt"
	"os"
)

func main() {
	var str string
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	fmt.Println(str)
}