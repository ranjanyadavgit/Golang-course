package main

import "fmt"

func main() {

	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println("length of array is :", len(array1))
	fmt.Println("length of array is:", len(array2))
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 10main.go
length of array is : 3
length of array is: 5
