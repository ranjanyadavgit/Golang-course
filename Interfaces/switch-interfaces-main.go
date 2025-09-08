package main

import (
	"fmt"
)

func main() {

	display(123)
	display("Go programming")
	display(15.64)
	display(true)
}

func display(a interface{}) {

	switch a.(type) {

	case int:
		fmt.Printf("Type:%T -- Value:%v\n", a, a.(int))

	case string:
		fmt.Printf("Type:%T -- Value:%v\n", a, a.(string))

	case float64:
		fmt.Printf("Type:%T -- Value:%v\n", a, a.(float64))

	default:
		fmt.Printf("type not found.")
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Interfaces
$ go run switch-interfaces-main.go 
Type:int -- Value:123
Type:string -- Value:Go programming
Type:float64 -- Value:15.64
type not found.
