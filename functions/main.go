package main

import(
  "fmt"
)

func testMultipleReturnValues()(int,int){
    return 3,4
}

func main(){
  fmt.Println("lets see some functions")
  fmt.Println("output with usual function :", vanillaFunctionToAdd(1,2))
  a,b := testMultipleReturnValues()
  fmt.Println("output with multiple return values are : ",a,b)
  _,c := testMultipleReturnValues()
  fmt.Println("output with multiple return values not caring about the first is  : ",c)

  array := []int{1,2,3,4,5}
  result := testVariadicFunctionToGetSumOfAllArray(array...)
  fmt.Println("sum of all element in array is ",result)
}


func vanillaFunctionToAdd(a int,b int) int{
  return a+b
}

func testVariadicFunctionToGetSumOfAllArray(nums ...int) int {
  sum := 0
  for _, element :=  range nums {
    sum+= element
  }
  return sum
}
