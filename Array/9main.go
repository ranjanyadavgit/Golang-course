package main

import "fmt"

func main() {

	array1 := [3]string{"alex", "jim", "bill"}

	array2 := array1
	array3 := &array1

	fmt.Printf("array1:%v\n", array1)
	fmt.Printf("array2:%v\n", array2)

	array1[0] = "ranjan"

	fmt.Printf("array1:%v\n", array1)
	fmt.Printf("array2:%v\n", array2)
	fmt.Printf("array3:%v", *array3)

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 7main.go
array1:[alex jim bill]
array2:[alex jim bill]
array1:[ranjan jim bill]
array2:[alex jim bill]
array3:[ranjan jim bill]
