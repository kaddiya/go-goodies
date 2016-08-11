package main

import(
  "fmt"
  alias_m "github.com/kaddiya/go-goodies/structs/models"
)

type Person struct{
  name string
  age int
}
func main(){
  fmt.Println("lets see some structs in action")
   a := alias_m.Employee{1}
  fmt.Println(alias_m.OuterPerson{"ameya",20,true})
  fmt.Println(a)
}
