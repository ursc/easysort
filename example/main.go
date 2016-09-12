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
	z := []User{User{"usera", 45}, User{"userc", 30}, User{"userb", 50}}

	fmt.Println("[IndexBy.Age]")
	indexByAge := easysort.IndexBy(len(z), func(i, j int) bool {
		return z[i].Age < z[j].Age
	})
	for _, i := range indexByAge {
		fmt.Printf("%+v\n", z[i])
	}

	fmt.Println("[OrderBy.Name]")
	easysort.OrderBy(len(z), func(i, j int) { z[i], z[j] = z[j], z[i] },
		func(i, j int) bool { return z[i].Name < z[j].Name },
	)
	for i := range z {
		fmt.Printf("%+v\n", z[i])
	}
}
