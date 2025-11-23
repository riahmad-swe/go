package main

import "fmt"

//* Print Welcome Message
func printWelcomeMessage()  {
	fmt.Println("Welcome to my Fokira Application")
}

//* Get User Name as input
func getUserName() string {
	var name string
	fmt.Println("Enter Your Name")
	fmt.Scanln(&name)
	return name
}

//* Get number input from user
func getTwoNumbers() (int, int) {
	var num1, num2 int
	fmt.Println("Enter First Number- ")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second Number- ")
	fmt.Scanln(&num2)
	return num1, num2
}

//* Summation of two numbers
func add(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

//* Display the result
func displayTheResult(name string, sum int)  {
	fmt.Println("Hello, ", name)
	fmt.Println("Summation of your input is- ", sum)
}

//* Print Goodbye Message
func printGoodbyeMessage()  {
	fmt.Println("Thank you for using my application")
	fmt.Println("Goodbye")
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	displayTheResult(name, sum)
	printGoodbyeMessage()


	//! Built a Summation app
	// //* Print Welcome Message
	// fmt.Println("Welcome to the My Fokira Application")

	// //* Get user name as input
	// var name string
	// fmt.Println("Enter Your Name")
	// fmt.Scanln(&name)

	// //* Get number input from user
	// var num1 int
	// var num2 int
	// fmt.Println("Enter First Number- ")
	// fmt.Scanln(&num1)
	// fmt.Println("Enter Second Number- ")
	// fmt.Scanln(&num2)

	// sum := num1 + num2

	// //* Display results
	// fmt.Println("Hello, ", name)
	// fmt.Println("Summation of your input is- ", sum)

	// //* Print a Goodbye Message
	// fmt.Println("Thank you for using my application")
	// fmt.Println("Goodbye")
}