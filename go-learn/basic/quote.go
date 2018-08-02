/*
写一个程序用来打印值类型和引用类型变量到终端 并观察输出结果
*/
package main

import (
"fmt"
"os"
)

func main() {
	var a = 100
	var b chan int = make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=".chan)
}
