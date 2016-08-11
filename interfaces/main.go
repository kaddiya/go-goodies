package main

import(
  "fmt"
  alias_models "github.com/kaddiya/go-goodies/interfaces/models"
)

func main(){
  fmt.Println("lets see some interfaces in actions")
  zoo := []alias_models.Animal{alias_models.Dog{},alias_models.Cat{},alias_models.Wolf{}}

  for _,animalx := range zoo{
    fmt.Println(animalx.Express())
  }
}
