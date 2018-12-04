package main
/*
	用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。
*/
func  popCount2(x uint)  int{
	var count,bit int
	for x > 0 {
		bit = int(x) % 2
		count += bit
		x = x / 2
	}
	return count
}
