package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 3 {
			break
		}
	}
}
Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/loops
$ go run 4main.go
0
1
2
3
