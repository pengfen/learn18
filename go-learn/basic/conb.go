/*
需求 编写代码 对于给定一个数字n 求出所有两两相加等于n的组合

如 对于a=2 所有组合如下
0 + 2 = 2
1 + 1 = 2
2 + 0 = 2
*/
package main

import "fmt"

func main() {
	list(1)
}

func list(n int) {
	for a := 0; a <= n; a++ {
		var b int = n - a // b declared and not used 注意var定义的值得使用，不然报错
		fmt.Printf("%d + %d = %d\n", a, b, n)
	}
}