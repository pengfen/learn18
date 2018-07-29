/*
使用包别名来访问包中的函数
*/

package main

import (
a "add"
"fmt"
)

func main() {
	
	fmt.Println("result: ", a.Name)
	fmt.Println("result: ", a.Age)
}