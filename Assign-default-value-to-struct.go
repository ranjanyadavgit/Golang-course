package main

import "fmt"

type Student struct {

	//student struct
	Name string
	Age  int
}

// instant method
func (s *Student) Instant() {

	if s.Name == "" {
		s.Name = "ranjan"
	}

	if s.Age == 0 {
		s.Age = 25
	}
}

func main() {

	student1 := Student{}
	student1.Instant()

	fmt.Println("student1:", student1)

	student2 := Student{
		Name: "julia",
	}
	student2.Instant()
	fmt.Println(student2)

	student3 := Student{
		Age: 30,
	}
	student3.Instant()

	fmt.Println(student3)

}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/struct
$ go run 7main.go 
student1: {ranjan 25}
{julia 25}
{ranjan 30}

