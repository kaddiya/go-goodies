package main

import(
  "fmt"
  alias_m "github.com/kaddiya/go-goodies/structs/models"
)

func main(){
  fmt.Println("lets see some structs in action")
  // employee := alias_m.Employee{1}
  outer_person := alias_m.OuterPerson{"ameya-2",20,true}
  outer_person.Talk()
}
