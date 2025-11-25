package main

import "fmt"

//! Slice has 3 values: 1. pointer, 2. lengh, 3. capacity
//? 1. Pointer: slice korar somoy je index theke count start hoi sei index tai pointer
//* 2. lengh: koto tomo index theke koto tomo index er age porjonto slice korbo seitai hochche slice er lengh. [1:4] = meaning index 1 theke index 3 porjonto.
//? 3. capacity: je index theke slice start hobe sekhan theke array er last index porjonto total index hochche slice er capacity. slice koto index korchi seita matter kore na. array er last index porjonto index count hobe capacity bujhaite.


func main() {
	arr := [6] string{"This", "is", "a", "must", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s)
}