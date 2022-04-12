package main

import (
	"fmt"
	"os"
)

var numbers = []int{1, 4, 3, 3, 2}
var target = 6

func main() {
	m := make(map[int]int)

	for i, v := range numbers {
		secIndex, ok := m[v]
		if ok {
			fmt.Printf("%v and %v", i, secIndex)
			return
		}

		c := target - v
		m[c] = i
	}

	fmt.Printf("found nothing")
	os.Exit(1)

}
