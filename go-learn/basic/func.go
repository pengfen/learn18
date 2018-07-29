package main

import "fmt"

func main() {
	// 定义局部变量
	var a int = 100
	var b int = 200
	var ret int

	// 调用函数并返回最大值 
	ret = max(a, b)
	fmt.Printf("最大值是: %d\n", ret)

    // 
	c, d := swap("ricky", "caopeng")
	fmt.Println(c, d)
}

// 返回多个值
func swap(x,y string) (string, string) {
	return y, x
}

func max(num1, num2 int) int {
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}