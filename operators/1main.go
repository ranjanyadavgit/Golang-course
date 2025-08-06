package main

import "fmt"

func main() {

	a := 50
	b := 30
	c := 20
	result1 := a + b
	result2 := a - b
	result3 := a * b
	result4 := a / b
	result5 := a % b
	fmt.Printf("the addition result of a + b = %d\n", result1)
	fmt.Printf("the substraction result of a - b = %d\n", result2)
	fmt.Printf("the multiplication result of a * b = %d\n", result3)
	fmt.Printf("the divison result of a / b = %d\n", result4)
	fmt.Printf("the modulous result of a %% b = %d\n", result5)

	fmt.Printf("before increment: %d\n", c)
	c++
	fmt.Printf("after increment: %d\n ", c)

}



o/p -

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/operators
$ go run main.go
the addition result of a + b = 80
the substraction result of a - b = 20
the multiplication result of a * b = 1500
the divison result of a / b = 1
the modulous result of a % b = 20
before increment: 20
after increment: 21


