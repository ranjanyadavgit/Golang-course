package main

import "fmt"

func swap(x, y int) int {

	var temp int

	temp = x

	x = y
	y = temp
	return temp

}

func main() {

	var a int = 100
	var b int = 200

	fmt.Printf("Before Swap, A = %d and B = %d\n", a, b)
	swap(a, b)
	fmt.Printf("After Swap, A = %d and B = %d", a, b)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 6main.go
Before Swap, A = 100 and B = 200
After Swap, A = 100 and B = 200
