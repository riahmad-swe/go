package main

import "fmt"

const name = "Golang"

func main() {
	const(
		port = 5000
		host = "localhost"
		userName = "riahmad"
		password = 123456
	)

	fmt.Println(port, host, userName, password)
}