# easysort

Easy sort in Go

## Example

```go

import (
	"fmt"
	"github.com/ursc/easysort"
)

type User struct {
	Name string
	Age  int
}

func main() {
	z := []User{User{"user1a", 45}, User{"user11c", 30}, User{"user2b", 50}}
	
	byID := easysort.IndexBy(len(z), func(i, j int) bool {
		return z[i].Age < z[j].Age
	})

	easysort.OrderBy(len(z), func(i, j int) { z[i], z[j] = z[j], z[i] },
		func(i, j int) bool { return z[i].Name < z[j].Name },
	)

	easysort.OrderBy(len(z), func(i, j int) { z[i], z[j] = z[j], z[i] },
		func(i, j int) bool { return easysort.NaturalLess(z[i].Name, z[j].Name) },
	)
}
```

Added [natural compares](https://github.com/fvbommel/util/blob/master/sortorder/natsort.go) two strings, using the natural order.
