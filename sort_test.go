package easysort

import (
	"testing"
)

type User struct {
	Name string
	Age  int
}

var testA = User{"usera", 45}
var testB = User{"userc", 30}
var testC = User{"userb", 50}

func TestIndexBy(t *testing.T) {
	z := []User{testA, testB, testC}

	idxByAge := IndexBy(len(z), func(i, j int) bool {
		return z[i].Age < z[j].Age
	})

	a := []int{1, 0, 2}

	for i, v := range idxByAge {
		if a[i] != v {
			t.Errorf("      data %+v", z)
			t.Errorf("indexByAge %+v", idxByAge)
			t.Errorf("       got %+v", a)
			return
		}
	}
}

func TestOrderBy(t *testing.T) {
	z := []User{testA, testB, testC}

	OrderBy(len(z), func(i, j int) { z[i], z[j] = z[j], z[i] },
		func(i, j int) bool { return z[i].Name < z[j].Name },
	)

	a := []User{testA, testC, testB}

	for i := range z {
		if a[i] != z[i] {
			t.Errorf("orderByName %+v", z)
			t.Errorf("        got %+v", a)
			return
		}
	}
}
