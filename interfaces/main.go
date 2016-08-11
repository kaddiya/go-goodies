package main

import(
  "fmt"
  alias_models "github.com/kaddiya/go-goodies/interfaces/models"
)

func main(){
  fmt.Println("lets see some interfaces in actions")
  zoo := []alias_models.Animal{alias_models.Dog{},alias_models.Cat{},alias_models.Wolf{},alias_models.Llama{}}

  for _,animalx := range zoo{
    value,err := animalx.Express()
    if(err == nil){
        fmt.Println(value)
    }else{
        fmt.Println("failed b/c:",err)
    }

  }
}
