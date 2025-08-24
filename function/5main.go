package main

import "fmt"

func rectangle(w int, h int) (int, int) {

	perimeter := 2 * (w + h)
	area := w * h

	return area, perimeter

}

func main() {

	var a, b int

	a, b = rectangle(20, 30)

	fmt.Println("area:", a)
	fmt.Println("perimeter:", b)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 5main.go
area: 600
perimeter: 100
