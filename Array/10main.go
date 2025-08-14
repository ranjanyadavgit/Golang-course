package main

import "fmt"

func main() {

	var array1 [3][3]int

	array1[0][0] = 1
	array1[0][1] = 4
	array1[0][2] = 6
	array1[1][0] = 6
	array1[1][1] = 9
	array1[1][2] = 4
	array1[2][0] = 12
	array1[2][1] = 2
	array1[2][2] = 8

	fmt.Println("elements are:")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", array1[i][j])
		}
		fmt.Println("")
	}

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 9main.go
elements are:
1       4       6
6       9       4
12      2       8

