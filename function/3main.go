package main

import "fmt"

func add(x int, y int) int {

	total := x + y
	return total
}

func main() {
	sum := add(10, 20)
	fmt.Printf("sum of X+Y = %d", sum)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 2main.go
total of X+Y = 11
Lenovo@DESKTOP-M5DT73G MINGW
