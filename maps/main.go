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
  fmt.Println("going to delete the value with key 1 from map_of_int_keys_and_string_values ")
  delete(map_of_int_keys_and_string_values,1)
  fmt.Println(map_of_int_keys_and_string_values)

  fmt.Println("going to check for existence  a non existing key")
  _,is_missing_key_present := map_of_int_keys_and_string_values[3]
  fmt.Println(is_missing_key_present)

  fmt.Println("going to check for existence an existing key")
  _, is_available_key_present := map_of_int_keys_and_string_values[2]
  fmt.Println(is_available_key_present)

  fmt.Println("going to dereference an existing key")
  exiting_value := map_of_int_keys_and_string_values[2]
  fmt.Println(exiting_value)

  fmt.Println("going to dereference a non-existing key")
  non_exiting_value := map_of_int_keys_and_string_values[3]
  fmt.Println(non_exiting_value)


}
