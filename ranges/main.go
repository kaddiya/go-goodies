package main

import (
	"fmt"
)

func main() {
	fmt.Println("lets see some ranges")
	nums := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, element := range nums {
		sum += element
	}

	println("sum of all elements using range is :", sum)
}
