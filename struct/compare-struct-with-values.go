package main

import "fmt"

//student struct

type Student struct {
	Name string
	Age  int
}

func main() {

	student1 := Student{
		Name: "Ranjan",
		Age:  29,
	}
	student2 := Student{
		Name: "Ranjan",
		Age:  29,
	}

	if student1 == student2 {
		fmt.Println("student1 variable is equal to student2 variable")
	} else {
		fmt.Println("student1 variable is not equal to student2 variable")
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/struct
$ go run 8main.go 
student1 variable is equal to student2 variable
