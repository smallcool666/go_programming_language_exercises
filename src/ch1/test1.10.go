package  main
/*
	 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，
修改本节中的程序，将响应结果输出，以便于进行对比。
*/
import "fmt"

func main()  {
	/*
	执行 go run fetchall.go https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
							https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
							https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
 	（这个url是go程序设计语言一书的一个中文翻译项目中的一节）
	得到输出如下：
1.88s    65412  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
1.90s    65407  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
2.08s    65413  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
第二次执行
1.77s    65412  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
2.50s    65413  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
2.63s    65413  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
第三次
1.76s    65406  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
1.81s    65413  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
1.83s    65413  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
第四次
1.87s    65407  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
2.11s    65413  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
3.00s    65406  https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch1/ch1-06.md
	*/
	fmt.Println("")
}