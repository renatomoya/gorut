Gorut
=====

## How to use?

```bash
go get github.com/renatomoya/gorut
```

```go
package main

import (
  "fmt"
  "github.com/renatomoya/gorut"
)

func main() {
  ok, err := gorut.ValidateRut("14696787-6")

  if err != nil {
    fmt.Printf("%v\n", err)
  }

  fmt.Printf("%v\n", ok) // true

  // What do you get when it's an invalid RUT?
  ok, err = gorut.ValidateRut("14696787-5")

  if err != nil {
    fmt.Printf("%v\n", err) // RUT is invalid
  }

  rut := gorut.Rut{Numbers: "14696787", Digit: "6"}
  ok, err = rut.IsValid()

  if err != nil {
    fmt.Printf("%v\n", err)
  }

  fmt.Printf("%v\n", ok)           // true
  fmt.Printf("%v\n", rut.Format()) // 14.696.787-6
}
```
