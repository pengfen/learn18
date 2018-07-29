/*
定义两个常量Man=1和Female=2 获取当前时间秒数 如果参被Female整除 则在终端打印female否则打印man
*/
package main

import (
"fmt"
"time"
)

func main() {
	const Man int64 = 1
	const female int64 = 2

	Second := time.Now().Unix() // 获取当前时间戳
	if (Second % Man == 0) {
		fmt.Print("female")
	} else {
		fmt.Print("male")
	}
}