package main

import "fmt"

//! What is Scope?= program er je kono ongshe kono variable ke access korte parle bujha jai Scope ache na parle Scope nai. Scope are 3 kinds: 1. Global Scope, 2. Local Scope 3. Package Scope

//? Global Scope
var a, b = 20, 30

func add(x int, y int)  {
	z := x + y
	fmt.Println(z)
}

func main() {
	var p = 30
	var q = 40

	add(p, q)

	add(a, b)

	add(a, p)
	
	add(b, q)

	// add(a, z)
}