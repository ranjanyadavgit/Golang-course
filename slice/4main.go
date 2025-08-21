package main

import "fmt"

func main() {

	var slice1 = make([]int, 4, 7)
	fmt.Printf("Slice1 = %v,\t length = %d, \tCapacity = %d\n", slice1, len(slice1), cap(slice1))

	var slice2 = make([]int, 7)
	fmt.Printf("Slice2 = %v,\tlength = %d, \tCapacity = %d\n", slice2, len(slice2), cap(slice2))
}



Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/slice
$ go run 4main.go
Slice1 = [0 0 0 0],      length = 4,    Capacity = 7
Slice2 = [0 0 0 0 0 0 0],       length = 7,     Capacity = 7

