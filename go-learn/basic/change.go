/*
写一个程序 交换两个整数的值
a=3;b=4 交换之后 a=4;b=3
*/

package main

import "fmt"

func main() {
	var a int = 3
	var b int = 4

	fmt.Printf("a=%d;b=%d\n", a, b);

	change(a, b)
}

func change(a int, b int) {
	var c int = a
	a = b
	b = c
	fmt.Printf("a=%d;b=%d\n", a, b);
}