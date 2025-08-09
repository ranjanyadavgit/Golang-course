package main

import "fmt"

func main() {

	names := [5]string{"Ranjan", "Jim", "Alex", "David"}

	fmt.Println("elements of arrays are")

	for i := 0; i < 4; i++ {
		fmt.Println("Names of array are", names[i])
	}
}
----

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 2main.go
elements of arrays are
Names of array are Ranjan
Names of array are Jim
Names of array are Alex
Names of array are David
