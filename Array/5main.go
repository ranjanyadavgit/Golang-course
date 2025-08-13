package main

import "fmt"

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(intArray); i++ {

		fmt.Println(intArray[i])
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 5main.go
1
2
3
4
5
