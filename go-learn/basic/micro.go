package main

import (
   "fmt"
   "time"
)

func main() {
	//记录开始时间
	//start := time.Nanoseconds()
	start := time.Now()
 
	//计算过程
	sum := 0
	for i := 0; i <= 100000000; i++{
		sum += i
	}
 
	//记录结束时间
	//end := time.Nanoseconds()
 
	//输出执行时间 单位为微秒
	//fmt.Println((end - start) / 1000000000)
 
	//输出执行结果
	//fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}