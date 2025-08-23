package main

import "fmt"

func main() {

	var employee = map[int]string{100: "Jim", 200: "kim", 300: "ryan"}

	fmt.Println("value of employee is:", employee[100])
	fmt.Println(employee[200])

}


$ go run 4main.go 
value of employee is: Jim
kim
