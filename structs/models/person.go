package models

import(
  "fmt"
)

type OuterPerson struct{
  Name string
  Age int
  IsAdmin bool
}

func (ptr *OuterPerson) Talk() {
  fmt.Println("Hello,I am "+ptr.Name)
}
