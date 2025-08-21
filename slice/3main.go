package main

import "fmt"

func main() {

	mainslice := [7]int{10, 20, 30, 40, 50, 60, 70}

	var slice1 = mainslice[1:5]

	slice2 := mainslice[0:]
	slice3 := mainslice[:6]
	slice4 := mainslice[:]
	slice5 := mainslice[2:4]

	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/slice
$ go run 3main.go 
slice1:  [20 30 40 50]
slice2: [10 20 30 40 50 60 70]
slice3: [10 20 30 40 50 60]
slice4: [10 20 30 40 50 60 70]
slice5: [30 40]
