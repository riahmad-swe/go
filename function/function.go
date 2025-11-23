package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return  sum
}

func getNumbers(num1 int, num2 int) (int, int)  {
	sum := num1 + num2
	multi := num1 * num2
	return sum, multi
}

func printSomething()  {
	fmt.Println("Education must be free!")
}

func sayHello(name string)  {
	fmt.Println("Welcome to the Golang course, ", name)
}

func main() {
	
	a := 10
	b := 20

	add(a, b)

	add(20, 30)

	//! Function with return

	sum := add(a, b)

	fmt.Println(sum)
	
	p, q := getNumbers(a, b)
	fmt.Println(p)
	fmt.Println(q)

	printSomething()
	sayHello("Ahmad")
}