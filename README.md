# ccpa

A Golang implementation of the IAB U.S. Privacy String (CCPA Opt-Out Storage Format)

You can read the full spec [here](https://iabtechlab.com/wp-content/uploads/2019/11/U.S.-Privacy-String-v1.0-IAB-Tech-Lab.pdf).

To install:

```
go get -v github.com/LiveRamp/ccpa
```

This package defines a struct (`Consent`) which contains all of the fields of the IAB Consent String.
The function Parse(s string) accepts the value of the `us_consent` parameter from the spec and returns
a `Consent` with all relevant fields populated (or an error if one occurred).

Example use:

```
package main

import (
  "fmt"

  "github.com/LiveRamp/ccpa"
)

func main() {
  var c, err = ccpa.Parse("1YNN")
  if err != nil {
    panic(err)
  }
  fmt.Printf("%+v\n", c)
}
```
