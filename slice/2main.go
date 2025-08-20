package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	slice1 := array[1:2]

	slice2 := array[0:]

	slice3 := array[:2]

	slice4 := array[:]

	fmt.Println("Array:", array)

	fmt.Println("Slice:", slice1)

	fmt.Println("slice2:", slice2)

	fmt.Println("slice3", slice3)

	fmt.Println("slice4", slice4)
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/slice
$ go run 2main.go
Array: [1 2 3 4 5]
Slice: [2]
slice2: [1 2 3 4 5]
slice3 [1 2]
slice4 [1 2 3 4 5]
