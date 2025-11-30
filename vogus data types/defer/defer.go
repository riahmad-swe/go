package main

import "fmt"

/*
	Rules of Named Return Values:
	1. all codes execute
	2. defer function store kora hobe
	3. return -> all defer functions execute after other execution
	4. return korbe named variable gulor values

	Rules of just return types that means un-named return values:
	1. all codes execute
	2. derer function store kora hobe
	3. return values are evaluated or stored at this time
	4. all defer functions execute korbe
*/

//! Named return value
func calc() (result int)  {
	fmt.Println("first: ", result)
	show := func ()  {
		result = result + 10
		fmt.Println("defer: ", result)
	}
	defer show()

	result = 5
	fmt.Println("second: ", result)

	p := func (a int)  {
		fmt.Println("ami: ", a)
	}
	defer p(result)

	defer fmt.Println(result)

	fmt.Println("third: ", result)
	
	defer fmt.Println(5)
	return
}
func calculate() int  {
	result := 0
	fmt.Println("first: ", result)
	show := func ()  {
		result = result + 10
		fmt.Println("defer: ", result)
	}
	defer show()

	result = 5
	fmt.Println("second: ", result)
	return result
}



// func a()  {
// 	i := 0

// 	fmt.Println("first: ", i)

// 	defer fmt.Println("second: ", i)

// 	i++

// 	fmt.Println("third: ", i)

// 	defer fmt.Println("forth: ", i)
	
// }

// func sum(a int, b int) (result int) { //? parameter a variable nile r pore variable declare korte hobe na.
// 	result = a + b
// 	return result
// }

func main() {
	a := calc()
	fmt.Println("calc: ", a)
	b := calculate()
	fmt.Println("calculate: ", b)
	// a()
	// res := sum(3, 4)
	// fmt.Println(res)
}