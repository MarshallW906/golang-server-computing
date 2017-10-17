package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5}
	for _, a := range array[1:] {
		fmt.Printf("%v", a)
	}
}
