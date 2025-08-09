package main

import "fmt"

func main() {

	var name [3]string

	name[0] = "Ranjan"
	name[1] = "Yadav"
	name[2] = "Jim"

	fmt.Println("elements of array")

	fmt.Println("first elements: ", name[0])
	fmt.Println("second elements:", name[1])

}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run main.go
elements of array
first elements:  Ranjan
second elements: Yadav
