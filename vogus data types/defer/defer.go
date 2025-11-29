package main

import "fmt"

func a()  {
	i := 0

	fmt.Println("first: ", i)

	defer fmt.Println("second: ", i)

	i++

	fmt.Println("third: ", i)

	defer fmt.Println("forth: ", i)
	return
}

func sum(a int, b int) (result int) { //? parameter a variable nile r pore variable declare korte hobe na.
	result = a + b
	return result
}

func main() {
	a()
	res := sum(3, 4)
	fmt.Println(res)
}