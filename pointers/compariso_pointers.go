package main

import "fmt"

func main() {

	value1 := 100
	value2 := 200

	p1 := &value1
	p2 := &value2
	p3 := &value1

	result1 := p1 == p2

	fmt.Println("the result1: p1 equals p2", result1)

	result2 := p1 == p3

	fmt.Println("the result: p1 equuls p3", result2)

	result3 := p2 == p3

	fmt.Println("the result: p2 equals p3", result3)

	result4 := &p1 == &p3

	fmt.Println("the result: p1 equals p3", result4)

	fmt.Println("the result: p2 equals p3", result5)

	result6 := &p1 == &p2

	fmt.Println("the result: p1 equals p2", result6)
}



Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/pointers
$ go run comparison_pointer.go 
the result1: p1 equals p2 false
the result: p1 equuls p3 true
the result: p2 equals p3 false
the result: p1 equals p3 false
the result: p2 equals p3 false
the result: p1 equals p2 false
