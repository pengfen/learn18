package main

import (
"fmt"
"strings"
)

func main() {
	var s  string = "http://www.res.com"
	var pre string = "http://"
	if (strings.HasPrefix(s, pre)) {
		fmt.Println("包括")
	} else {
		fmt.Println("不包括")
	}

	var suf string = "/"
	if (strings.HasSuffix(s, suf)) {
		fmt.Println("包括")
	} else {
		fmt.Println("不包括")
	}

	var index string = "w"
	fmt.Println(strings.Index(s, index)) // 7
	fmt.Println(strings.LastIndex(s, index)) // 9

	fmt.Println(strings.Replace(s, index, "e", 1)) //
	fmt.Println(strings.Count(s, index))
	fmt.Println(s + strings.Repeat("a", 2))

}