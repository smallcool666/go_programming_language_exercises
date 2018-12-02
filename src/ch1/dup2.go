package main
/*
	书中的示例代码，用于统计重复的行
	如果有通过命令行传递文件名参数，就对每个文件的行进行统计，没有传递参数的话就从标准输入读入行内容并统计
	注：windows下Ctrl+d 完成输入，linux下Ctrl+c完成输入
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if (len(files) < 1) {
		countLines(os.Stdin, counts)
	} else {
		for _,v := range files {
			f, err := os.Open(v)
			if (err != nil) {
				fmt.Println("open File" + v + " error!")
				continue
			}
			countLines(f, counts)
		}
	}
	for k, v := range counts {
		fmt.Println(k, "  ", v)
	}
}

func countLines(f *os.File, c map[string]int) {
	input := bufio.NewScanner(f)
	for  input.Scan() {
		c[input.Text()]++
	}
}