package main

import "fmt"

func main() {
	//? Annonymous function
	//? Immediately Invoked Function Expression
	//? IIFE - that means which function has no name
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(4, 7) //! IIFE te immediately value declare kore call korte hobe. na hole error dibe.
}

func init()  {
	fmt.Println("I will be executed or invoked of executed first")
}