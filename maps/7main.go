package main

import (
	"fmt"
	"sort"
)

func main() {

	unsortedvalue := map[int]string{

		20: "ranjan",
		10: "yadav",
		50: "kumar",
		30: "electronic",
		40: "city",
	}

	keys := make([]int, 0, len(unsortedvalue))

	for i := range unsortedvalue {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for _, value := range keys {
		fmt.Println(value, unsortedvalue[value])
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/maps
$ go run 7main.go 
10 yadav
20 ranjan
30 electronic
40 city
50 kumar
