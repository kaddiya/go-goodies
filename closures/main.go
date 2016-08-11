package main

import (
	"fmt"
)

func main() {
	bar := foo()
	a := bar()
	fmt.Println(a)
	b := bar()
	fmt.Println(b)

	new_bar := foo()
	new_a := new_bar()
	fmt.Println(new_a)
	new_b := new_bar()
	fmt.Println(new_b)

}

func foo() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
