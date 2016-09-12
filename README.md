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
	z := []User{ User{"testa", 45}, User{"testb", 30} }
	
	byID := easysort.IndexBy(len(z), func(i, j int) bool {
		return z[i].Age < z[j].Age
	})

	easysort.OrderBy(len(z), func(i, j int) { z[i], z[j] = z[j], z[i] },
		func(i, j int) bool { return z[i].Name < z[j].Name },
	)
}
```