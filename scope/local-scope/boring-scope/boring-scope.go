package main

import "fmt"

var a, b = 20, 30

func printNum(num int)  {
	fmt.Println(num)
}

func add(x int, y int)  {
	result := x + y
	printNum(result)
}

func main() {
	add(a, b)
}