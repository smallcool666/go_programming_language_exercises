package main
/*
	修改dup2，使其可以分别打印重复的行出现在哪些文件。
	countLines方法再dup2中定义
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[int]string)
	files := os.Args[1:]
	if (len(files) < 1) {
		newCountLines(os.Stdin, counts, "stdin")
	} else {
		for _,v := range files {
			f, err := os.Open(v)
			if (err != nil) {
				fmt.Println("open File" + v + " error!")
				continue
			}
			newCountLines(f, counts, v)
		}
	}
	for k, item := range counts {
		//输出行内容
		fmt.Println(k)
		for _, v :=range item {
			fmt.Println("  ", v)
		}
	}
}

func newCountLines(f *os.File, c map[string]map[int]string, fileName string ){
	input := bufio.NewScanner(f)
	for  input.Scan() {
		if len(c[input.Text()]) == 0 {
			//二维map每一个元素都需要申请空间
			c[input.Text()] = make(map[int]string)
		}
		c[input.Text()][len(input.Text())] = fileName
	}
}