package models

type Cat struct{}
func (c Cat) Express() (string,error){
  return "meow!", nil
}
