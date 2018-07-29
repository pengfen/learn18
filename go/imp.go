/*
add包中有两个变量 Name和Age 访问这两个变量

注意变量首字母大写才可以访问
package add

var Name string = "welcome to go world"
var Age int = 10

*/
package main

import (
"add"
"fmt"
)

func main() {
	fmt.Println("result: ", add.Name)
	fmt.Println("result: ", add.Age)
}