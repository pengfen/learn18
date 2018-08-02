package main

import "fmt"

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	fmt.Println(a)
}

func m() {
	a := "o"
	fmt.Println(a)
}