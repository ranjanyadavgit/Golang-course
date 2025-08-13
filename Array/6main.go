package main

import "fmt"

func main() {

	intArray := [5]int{1, 2, 3, 4, 5}

	for index, value := range intArray {
		fmt.Println(index, "=>", value)
	}
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 6main.go
0 => 1
1 => 2
2 => 3
3 => 4
4 => 5
