package main

import (
	"fmt"
)

func main() {
	fmt.Println("lets see some slices")
	s := make([]int, 0)
	index := 0
	for index < 10 {
		s = append(s, index)
		index++
	}
	fmt.Println("slices is made of length : ", len(s))
	fmt.Println("slice that guy!", s[2:9])

}
