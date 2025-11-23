package main

import "fmt"


func add(a, b int)  {
	fmt.Println(a + b)
}


func main() {
	add(2, 3)
	add := func (a int, b int)  {
		c := a + b
		fmt.Println(c)
	}
	add(4, 7) //! annonymous function ke akta variable er moddhe rakhle sei function block er baire main function er vitore value diye call kora jai.
	
}

func init()  {
	fmt.Println("I'll be called first")
}