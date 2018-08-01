package main

import "fmt"

func main() {
	LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	i := 0
	HERE:
	    Print(i)
	    i ++
	    if i == 5 {
	    	return
	    }
	    goto HERE

	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("", i)
		i ++;
	}

	for i := 0; i <= 7; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}