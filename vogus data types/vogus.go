package main

import "fmt"

func main() {
	var a int8 = -128
	var b int8 = 127
	
	var x uint8 = 255 //? u mean unsigned but only positive values

	var j float32 = 10.23343
	var k float64 = 10.4455

	var flag bool = false
	var flag1 bool = true
	var s string = "My name is Ahmad"

	fmt.Println(a, b, x, j, k, flag, flag1)

	//! Print using f formatting
	fmt.Printf("%d\n", a) //? decimal -> "%d"
	fmt.Printf("%.2f\n", j) //? float -> "%f" . er pore digit define "%.3f"
	fmt. Printf("%v\n", flag) //? boolean -> "%v"
	fmt.Printf("%s\n", s) //? string -> "%s"


	//! rune
	r := '🔥'
	fmt.Println(r)
	fmt.Printf("%c", r) //? "%c" -> character print formatting (Printf)
}