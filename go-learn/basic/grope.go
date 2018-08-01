package main

import "fmt"

var a string


func main() {
	a = "G"
	fmt.Print(a)
	f1()
}

func f1() {
	a := "o"
	fmt.Println(a)
	f2()
}

func f2() {
	fmt.Println(a)
}