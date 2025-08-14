package main

import "fmt"

func main() {

	array1 := []string{"go", "java", "python", "c++", "c#"}

	fmt.Println("length:", len(array1))
	fmt.Println("prt the string:", array1)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 11main.go
length: 5
prt the string: [go java python c++ c#]
