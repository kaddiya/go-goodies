package models


type Dog struct{}
func (d Dog) Express() (string,error) {
  return "woof!",nil
}
