package main

import "fmt"

var a int = 25

func main() {
	fmt.Println(a)
}

func init() {
	fmt.Println(a)
	a = 50
}