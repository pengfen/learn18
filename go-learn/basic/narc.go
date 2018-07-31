package main

import "fmt"

func isSXH(n int) bool {
	/*
	打印出100-999中所有的水仙花数 所谓水仙花数是指一个三位数 其各位数字立方和等于该数本身
	*/
	var i, j, k int
	i = n % 10 // 获取个位数
	j = (n / 10) % 10
	k = (n / 100) % 10
	sum := i * i * i + j * j * j + k * k * k
	return sum == n
}

func main() {
	
}