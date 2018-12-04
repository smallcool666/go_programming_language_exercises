package main
/*
	表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。
*/
func  popCount3(x uint)  int{
	var count int
	for x > 0 {
		x = x & (x -1)
		count++
	}
	return count
}
