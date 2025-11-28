package main

import "fmt"

//! Slice has 3 values: 1. pointer, 2. lengh, 3. capacity
//? 1. Pointer: slice korar somoy je index theke count start hoi sei index tai pointer
//* 2. lengh: koto tomo index theke koto tomo index er age porjonto slice korbo seitai hochche slice er lengh. [1:4] = meaning index 1 theke index 3 porjonto.
//? 3. capacity: je index theke slice start hobe sekhan theke array er last index porjonto total index hochche slice er capacity. slice koto index korchi seita matter kore na. array er last index porjonto index count hobe capacity bujhaite.

/*
//* Types of Slice:
1. Slice from an existing array
2. slice from a slice
3. slice literal
4. make function with len
5. make function with len and cap
6. empty slice of nil slice
7. slice underlying array rule => 1024 -> double lengh will be multiply by 2. 1x2 2x2 3xnothing 4xnothing as like 1024. 1024 cross korle 25% kore increase korbe.
*/

//! Variadic Function using Slice
func print(numbers ...int)  {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func changedSlice(p [] int) [] int  {
	p[0] = 10
	p = append(p, 11)
	return p
}

func main() {
	print(5, 6, 7, 8, 9, 10)

	x := [] int {1, 2, 3, 4, 5, 6, 7,}
	x = append(x, 8)
	x = append(x, 9)

	a := x[4:]

	y := changedSlice(a)

	fmt.Println(x)
	fmt.Println(y)



	// var x [] int
	// x = append(x, 1)
	// x = append(x, 2)
	// x = append(x, 3)

	// y := x

	// x = append(x, 4)
	// y = append(y, 5)

	// x[0] = 10

	// fmt.Println(x) //? [10, 2, 3, 5]
	// fmt.Println(y) //? [10, 2, 3, 5]


	// s := [] int {1, 2, 5} //? Slice Literal (pointer, lengh, capacity)
	// fmt.Println("Slice: ", s, "Lengh: ", len(s), "Capacity: ", cap(s))

	//! Make slice using built in function
	// s := make([] int, 3) //* [0, 0, 0], len = 3, cap = 3

	// s := make([] int, 3, 5) //* [0, 0, 0], len = 3, cap = 5
	// s[0] = 5
	// s[2] = 2
	// s[3] = 7
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// var s [] int
	// fmt.Println(s)

	//! make slice using append() function

	// var s [] int // len = 0, cap = 0
	// s = append(s, 1)
	// s = append(s, 1, 2, 3)
	// fmt.Println(s)

	// arr := [6] string{"This", "is", "a", "must", "interview", "question"}
	// fmt.Println(arr)

	// s := arr[1:4]
	// fmt.Println(s)
}