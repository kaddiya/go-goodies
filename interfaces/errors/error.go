package errors

import(
  "fmt"
)
type ArgError struct {
    Arg  int
    Prob string
}
func (e *ArgError) Error() string {
    return fmt.Sprintf("%d - %s", e.Arg, e.Prob)
}
