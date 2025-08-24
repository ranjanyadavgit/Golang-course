package main

import "fmt"

func main() {

	fmt.Printf("Rectangle value is = %d\n", func(w int) int {
		return w * w

	}(10))
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 10main.go 
Rectangle value is = 100
