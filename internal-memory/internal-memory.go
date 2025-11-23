package main

import "fmt"

//! Memory management: Code Segment (all codes stay here) => Data segment (all global variable stays here) => Stack (all execution will happen here through Stack Frame) => Heap (GC means Garbage Collector)

var a = 10

func add(x, y int)  {
	z := x + y
	fmt.Println(z)
}

func main() {
	add(5, 4)
	add(a, 3)
}

func init()  {
	fmt.Println("Salam")
}