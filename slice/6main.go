package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70}

	for index, value := range slice {
		fmt.Printf("Index = %d and Value = %d\n", index, value)
	}

	for _, value := range slice {
		fmt.Printf("the value = %d\n", value)
	}

	i := 0
	for range slice {
		fmt.Println("", slice[i])
		i++
	}

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/slice
$ go run 5main.go
Index = 0 and Value = 10
Index = 1 and Value = 20
Index = 2 and Value = 30
Index = 3 and Value = 40
Index = 4 and Value = 50
Index = 5 and Value = 60
Index = 6 and Value = 70
the value = 10
the value = 20
the value = 30
the value = 40
the value = 50
the value = 60
the value = 70
 10
 20
 30
 40
 50
 60
 70

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/slice
$
