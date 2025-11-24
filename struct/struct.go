package main

import "fmt"

type User struct {
	Name string
	Age int
}

type Ahnaf struct{
	Name string
	Age int
	School string
	Class int
	Status string
}

func main() {

	var user1 User
	user1 = User{
		Name: "Habib",
		Age: 30,
	}
	fmt.Println("Name: ", user1.Name)
	fmt.Println("Age: ", user1.Age)

	user2 := User{
		Name: "Roki",
		Age: 16,
	}
	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age: ", user2.Age)

	var ahnaf Ahnaf
		ahnaf = Ahnaf{
		Name: "Muhammad Ahnaf",
		Age: 13,
		School: "Darus Salam Kamil Madrasha",
		Class: 6,
		Status: "The Professor is a Good Boy",
	}

	fmt.Println("Name: ", ahnaf.Name)
	fmt.Println("Age: ", ahnaf.Age)
	fmt.Println("School: ", ahnaf.School)
	fmt.Println("Class: ", ahnaf.Class)
	fmt.Println("Status: ", ahnaf.Status)

	arhab := Ahnaf{
		Name: "Muhammad Arhab",
		Age: 4,
		School: "Home Garden",
		Class: 1,
		Status: "Very Humble Boy",
	}
	fmt.Println("Name: ", arhab.Name)
	fmt.Println("Age: ", arhab.Age)
	fmt.Println("School: ", arhab.School)
	fmt.Println("Class: ", arhab.Class)
	fmt.Println("Status: ", arhab.Status)

	yusra := Ahnaf{
		Name: "Zunainah Binte Ahmad Yusra",
		Age: 2,
		School: "Home Garden",
		Class: 1,
		Status: "Very Naughty Girl",
	}

	fmt.Println("Name: ", yusra.Name)
	fmt.Println("Age: ", yusra.Age)
	fmt.Println("School: ", yusra.School)
	fmt.Println("Class: ", yusra.Class)
	fmt.Println("Status: ", yusra.Status)

}