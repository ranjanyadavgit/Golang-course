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
}
