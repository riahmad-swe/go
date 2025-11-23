package main

import "fmt"

//! Higher Order Function: / First Class Function
//? HOF works to set a rule inside a FOF. Rules are:
/*
	1. parameter -> function
	2. function return
	3. both
*/

// func processOperation(a int, b int, op func (p int, q int))  {
// 	op(a, b)
// }

//! Callback function (with return value)
func call() func (x int, y int) {
	return add
}
func add(x int, y int){
	z := x + y
	fmt.Println(z)
}

func main() {

	// processOperation(2, 5, add)

	sum := call() //* function assign kora ke bole Function Expression

	sum(4, 3)

	//! First Class Citizen => variable er moddhe je data assign kora hoi take first class citizen bole


	


}