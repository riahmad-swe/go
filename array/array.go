package main

import "fmt"

var arr2 = [3]string{	//! Shorthand array declaration
	"Ahnaf,", "Arhab,", "Yusra"}

func main() {
// var arr1 [2] int //! normal array declaration
// 	arr1[0] = 5
// 	arr1[1] = 7

	// fmt.Println(arr1)

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(arr2)
}
