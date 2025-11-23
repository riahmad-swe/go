package main

import "fmt"

func main() {
	//! Condition
	age := 18
	if age > 18 {
		fmt.Println("You are eligible to be married")
	} else if age < 18 {
		fmt.Println("You are not eligible to be married")
	} else if age == 18 {
		fmt.Println("You are just a teenager, eat a lolipop baby")
	}

	//! conditional operators: < > <= >= == && || !

	age1 := 20
	sex := "male"
	if age1 == 20 && sex == "male" {
		fmt.Println("You are eligible to be married")
	}
	if age > 40 || sex == "male" {
		fmt.Println("Ready, Go!")
	}

	isPretty := false
	if isPretty {
		fmt.Println("You are pretty")
	}

	a := 4
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2, 3:
		fmt.Println("a is either 2 or 3")
	default:
		fmt.Println("a is neither 1 or 2 or 3")
	}
	b := 10
	switch b {
		case 1, 2, 3, 4, 5:
			fmt.Println("1 to 5")
		case 6, 7, 8:
			fmt.Println("6 to 8")
		default:
			fmt.Println("other")
	}
}