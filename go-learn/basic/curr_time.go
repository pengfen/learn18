package main

import (
    "fmt"
    "time"
)

func main() {
	curr := time.Now()
	fmt.Printf("%2d/%2d/%2d %2d:%2d:%2d", curr.Year(), curr.Month(), curr.Day(), curr.Hour(), curr.Minute(), curr.Second())
}