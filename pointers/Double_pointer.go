package main

import "fmt"

func main() {

	var x int = 100

	var pt1 *int = &x

	var pt2 **int = &pt1

	fmt.Println("the value of x = ", x)

	fmt.Println("the value of PT1 = ", pt1)

	fmt.Println("the value of  PT2 = ", pt2)

	fmt.Println("value of address of pt2  or *pt2 =", *pt2)
	fmt.Println("value of address of pt2 or **pt2 =", **pt2)

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/pointers
$ go run Double_pointer.go 
the value of x =  100
the value of PT1 =  0xc00000a0e8
the value of  PT2 =  0xc000066060
value of address of pt2  or *pt2 = 0xc00000a0e8
value of address of pt2 or **pt2 = 100
