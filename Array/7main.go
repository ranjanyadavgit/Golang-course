package main

import "fmt"

func main() {

	intArray := [5]int{1, 2, 3, 4, 5}


	for _, value := range intArray {
		fmt.Println(value)
	}

}
o/p ->
1
2
3
4
5
