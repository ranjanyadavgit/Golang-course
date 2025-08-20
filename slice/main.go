package main

import "fmt"

func main() {

	var slice1 = []string{"go", "java", "c#", "perl"}

	fmt.Println("slices:", slice1)

	slices2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("values are:", slices2)
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/slice
$ go run main.go
slices: [go java c# perl]
values are: [1 2 3 4 5 6]
