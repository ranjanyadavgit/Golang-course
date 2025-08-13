package main

import "fmt"

func main() {

  intArray := [5]int{1, 2, 3, 4, 5}
  
  j := 0

	for range intArray {
		fmt.Println(intArray[j])
		j++
	}

}

o/p - >

1
2
3
4
5
