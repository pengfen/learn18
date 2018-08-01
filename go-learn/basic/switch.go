package main

import "fmt"

func main() {
	var i = 0
	switch i {
	    case 0:
	    case 1:
	    	fmt.Println(i)
	    case 2:
	    	fmt.Println(i)
	    default:
	    	fmt.Println("def")
	}

	switch i {
	    case 0:
	    	fallthrough
	    case 1:
	    	fmt.Println(i)
	    case 2:
	    	fmt.Println(i)
	    default:
	    	fmt.Println("def")
	}

	switch i {
	    case 0, 1:
	    	fmt.Println("1")
	    case 2:
	    	fmt.Println(i)
	    default:
	    	fmt.Println("def")
	}

	switch {
	    case i > 0 && i < 10:
	    	fmt.Println("i > 0 and i < 10")
	    case i > 10 && i < 20:
	    	fmt.Println("i > 10 and i < 20")
	    default:
	    	fmt.Println("def")
	}
}