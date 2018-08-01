package main

import "fmt"

func main() {
	str := "welcome to go world"

	for i, v := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(v)))
	}

	for i, v := range str {
		if i > 2 {
			continue
		}
		if (i > 3) {
			break
		}
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(v)))
	}
}
