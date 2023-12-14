package main

import "fmt"

func doubleNum(num int) int {
	return num * 2
}

func main() {
	x := 5
	fmt.Println(doubleNum(x))
}