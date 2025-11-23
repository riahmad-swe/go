package main

import (
	"fmt"

	"example.com/mathlibrary"
)

//! Local Scope means when we declare a variable in a block that calls a local scope.

var a, b int = 10, 20



func main() {
	// x := 18
	// p:= 25
	// if x>= 18 {
	// 	fmt.Println("I'm a matured boy")
	// 	fmt.Println("I've", p, "girlfriends")
	// }

	// add(4, 7)

	fmt.Println("Showing Custom Package")

	mathlibrary.Add(20, 30)
}