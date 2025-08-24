package main

import "fmt"

func add(x int, y int) {

	total := x + y
	fmt.Printf("total of X+Y = %d", total)
}

func main() {
	add(5, 6)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 2main.go
total of X+Y = 11
