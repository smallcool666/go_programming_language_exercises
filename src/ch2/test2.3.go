package main

func init() {
	//在每个文件中的init初始化函数，在程序开始执行时按照它们声明的顺序被自动调用。
	for i := range pc {
		//记录了0-255每个数字对应的1的个数
		pc[i] = pc[i/2] + byte(i & 1)
	}
}
func  popCount1(x uint)  int{
	var count int
	for x > 0 {
		count += int(pc[byte(x)])
		x = x >> 8
	}
	return count
}
