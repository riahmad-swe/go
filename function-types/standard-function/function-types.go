package main

import "fmt"

//! There are many types of functions.
//* 1. Standard of named function
//* Anonymous function
//* Function expression of Assign function in variable
//* Higher order function or first class function
//* Callback function
//* Variadic function
//* Init function - you can not call this, computer calls this automatically
//* Closure - close over
//* Defer function - last in first out
//* Receiver function
//* IIFE - Immediately invoked function expression

//! Standard or named function example:
func add(a, b int)  {   //? aikhane function er name ache 'add'
	fmt.Println(a + b)
}

func main() {
	add(4, 7)
}