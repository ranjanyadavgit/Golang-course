package main

import "fmt"

func main() {

	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/loops
$ go run 3main.go
1
2
3
4
