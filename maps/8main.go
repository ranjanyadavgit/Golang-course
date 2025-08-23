package main

import "fmt"

func main() {

	myMap := map[int]string{

		10: "ranjan",
		20: "Kumar",
		30: "Yadav",
	}

	for index, value := range myMap {
		fmt.Printf("index:%d\tValues:%s\n", index, value)
	}

	for i := range myMap {
		delete(myMap, i)

	}
	fmt.Println("my map after deleting")
	fmt.Println(myMap)
}

$ go run 8main.go 
index:10        Values:ranjan
index:20        Values:Kumar
index:30        Values:Yadav
my map after deleting
map[]
