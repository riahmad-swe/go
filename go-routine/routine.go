package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func add(a, b int) {
	fmt.Println(a + b)
}

func printHello(num int)  {
	fmt.Println("Salam Baba", num)
	
}
func main() {
	add(2, 4)
	var x int = 10
	fmt.Println("hello", x)

	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)
	
	go printHello(5)

	fmt.Println(a, " ", p)

	time.Sleep(5 * time.Second)
}