package main

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

	fmt.Println("[IndexBy.Age]")
	indexByAge := easysort.IndexBy(len(z), func(i, j int) bool {
		return z[i].Age < z[j].Age
	})
	for _, i := range indexByAge {
		fmt.Printf("%+v\n", z[i])
	}

	fmt.Println("\n[OrderBy.Name]")
	easysort.OrderBy(len(z), func(i, j int) { z[i], z[j] = z[j], z[i] },
		func(i, j int) bool { return z[i].Name < z[j].Name },
	)
	for i := range z {
		fmt.Printf("%+v\n", z[i])
	}

	fmt.Println("\n[OrderBy.Natural.Name]")
	easysort.OrderBy(len(z), func(i, j int) { z[i], z[j] = z[j], z[i] },
		func(i, j int) bool { return easysort.NaturalLess(z[i].Name, z[j].Name) },
	)
	for i := range z {
		fmt.Printf("%+v\n", z[i])
	}
}
