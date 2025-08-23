package main

import "fmt"

func main() {

	var employee = make(map[string]int)

	employee["hey"] = 3
	employee["there"] = 5
	employee["byethere"] = 6

	employlist := make(map[string]int)
	fmt.Println(len(employee))
	fmt.Println(len(employlist))

}

$ go run 3main.go
3
0
