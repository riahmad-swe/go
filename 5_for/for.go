package main

import "fmt"

// for -> only construct in go for looping

func main() {
	
	i := 1
	for i<=5 {
		// fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic for loop
	// for i := 1; i <=7;	i++ {
		// break
		// continue
	// 	if i == 2{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println(i)

	// range
	for i := range 7 {
		fmt.Println(i)
	}
}