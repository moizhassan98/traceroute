# traceroute
 
## Import

To import the module first you have to `go get` the module which can be done with:

```cmd
go get github.com/moizhassan98/traceroute

```

and then import it using

```go
import (
    "github.com/moizhassan98/traceroute"
)
```

## Usage

There are 2 main functions that you can use to get hops. `GetHops` and `GetHopsJSON`.
`GetHops` function return a struct of type `Trace`.
`Trace` struct structure is given below.

```go
// Route represents information about a single hop in the trace route
type Route struct {
	Hop     int
	Address string
	Time1   string
	Time2   string
	Time3   string
}

// Trace represents the trace route result
type Trace struct {
	Destination     string
	Routes          []Route
	traceSuccessful bool
}
```
`GetHopsJSON` function returns a json in form of `[]byte`. It is just a Marshaled Trace.


## Examples

### `GetHops` Example

```go
package main

import (
	"fmt"

	"github.com/moizhassan98/traceroute"
)

func main() {

	trace, err := traceroute.GetHops()
	if err != nil {
		fmt.Println("Error in getting Hops!")
	}
	for _, hop := range trace.Routes {
		fmt.Printf("#%v , %v, %v, %v IP: %v \n", hop.Hop, hop.Time1, hop.Time2, hop.Time3, hop.Address)
	}
}
```