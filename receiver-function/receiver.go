package main

import "fmt"

type Family struct {
	Name   string
	Age    int
	School string
	Class  int
	Status string
}

// func printMyFamily(family Ahnaf)  {
// 	fmt.Println("Name: ", family.Name)
// 	fmt.Println("Age: ", family.Age)
// 	fmt.Println("School: ", family.School)
// 	fmt.Println("Class: ", family.Class)
// 	fmt.Println("Status: ", family.Status)
// }

//! Receiver Function
func (family Family) printMyFamily()  {
	fmt.Println("Name: ", family.Name)
	fmt.Println("Age: ", family.Age)
	fmt.Println("School: ", family.School)
	fmt.Println("Class: ", family.Class)
	fmt.Println("Status: ", family.Status)
}


func main() {
	var ahnaf Family
	ahnaf = Family{
		Name:   "Muhammad Ahnaf",
		Age:    13,
		School: "Darus Salam Kamil Madrasha",
		Class:  6,
		Status: "The Professor is a Good Boy",
	}
	// printMyFamily(ahnaf)
	ahnaf.printMyFamily() 	//? receiver function call
	
	arhab := Family{
		Name:   "Muhammad Arhab",
		Age:    4,
		School: "Home Garden",
		Class:  1,
		Status: "Very Humble Boy",
	}
	// printMyFamily(arhab)
	arhab.printMyFamily()	//? receiver function call
	
	yusra := Family{
		Name:   "Zunainah Binte Ahmad Yusra",
		Age:    2,
		School: "Home Garden",
		Class:  1,
		Status: "Very Naughty Girl",
	}
	// printMyFamily(yusra)
	yusra.printMyFamily()	//? receiver function call
	
}