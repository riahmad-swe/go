package main

import "fmt"

func print(numbers *[10]int) {
	fmt.Println(numbers)
}

type YaHabibi struct{
	Name string
	Age int
	Salary float64
	FavFoods [] string
}

func main() {
	habib := YaHabibi{
		Name: "Habibur Rahman",
		Age: 30,
		Salary: 0,
		FavFoods: [] string {"Orange", "Malta", "Banana"},
		
	}

	p := &habib
	fmt.Println(*p)
	fmt.Println(habib.FavFoods)
	// fmt.Println(habib)

	// x := 20
	// p := &x	//! & hochche address or pointer of x.
	// fmt.Println("X =", x)
	// *p = 30
	// fmt.Println("X =", x)

	// 	fmt.Println("Address or Pointer is: ", p)
	// 	fmt.Println("Value of the Address: ", *p) //! * sign diye value of address bujhai
	// 	fmt.Println(p)

	// arr := [10] int{
	// 	1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// print(&arr)

	//? * = pointer or address of any kinds of values
	//? & = value of that particular pointer or address

	//! Pass by value
	//! Pass by reference ==== // pointer reference kei pass kore. puro array ba values gulo ke pass na kore just address dekhiye dei.
}