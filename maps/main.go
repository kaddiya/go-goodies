package main

import(
  "fmt"
)

func main(){
  map_of_int_keys_and_string_values := make(map[int]string)
  map_of_string_keys_and_int_values := make(map[string]int)
  fmt.Println("lets see some maps!")

  map_of_string_keys_and_int_values["first"] = 1
  map_of_string_keys_and_int_values["second"] = 2

  fmt.Println(map_of_string_keys_and_int_values)


  map_of_int_keys_and_string_values[1]= "first"
  map_of_int_keys_and_string_values[2]= "second"

  fmt.Println(map_of_int_keys_and_string_values)
}
