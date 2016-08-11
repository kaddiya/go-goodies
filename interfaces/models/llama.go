package models

import(
  err "github.com/kaddiya/go-goodies/interfaces/errors"
)

type Llama struct{
}

func (llama Llama) Express() (string,error) {
    return "???",&err.ArgError{422, "A Llama cant speak"}
}
