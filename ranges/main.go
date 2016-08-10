package main

import (
	"fmt"
)

func main() {
	fmt.Println("lets see some ranges")
  nums := make([]int, 5)

  nums[0] = 1
  nums[1] = 2
  nums[2] = 3
  nums[3] = 4
  nums[4] = 5
	sum := 0

	for _, element := range nums {
		sum += element
	}

  nums_squared := make([]int, 5)
  fmt.Println(len(nums_squared))
	fmt.Println("sum of all elements using range is :", sum)

  fmt.Println("now lets square every element of the num array")
  fmt.Println(len(nums_squared))
  for index,element := range nums{
      nums_squared[index] = element**2
  }
  fmt.Println(nums_squared)

}
